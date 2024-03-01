// Copyright (c) 2020-2024 Peter Hagelund
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/peterhagelund/go-irobot/irobot"
	"go.bug.st/serial"
)

func main() {
	mode := &serial.Mode{
		BaudRate: 115200,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
	port, err := serial.Open("/dev/ttyUSB0", mode)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected.")
	port.SetReadTimeout(time.Second * 1)
	roomba, err := irobot.NewRoomba(port)
	if err != nil {
		log.Fatal(err)
	}
	defer roomba.Close()
	if err = roomba.Start(); err != nil {
		log.Fatal(err)
	}
	defer roomba.Power()
	// Put the Roomba in safe mode so it can be driven (safely).
	if err = roomba.Safe(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Started.")
	defer roomba.DriveDirect(0, 0) // If things go south, at least try to stop the beast.
	// Go forward at 250 mm/s; then stop; then backwards at -250 mm/s; then stop.
	velocities := [][2]int16{{250, 250}, {0, 0}, {-250, -250}, {0, 0}}
	for i := range velocities {
		leftVelocity := velocities[i][0]
		rightVelocity := velocities[i][1]
		fmt.Printf("Driving %d, %d\n", leftVelocity, rightVelocity)
		if err := roomba.DriveDirect(leftVelocity, rightVelocity); err != nil {
			fmt.Println(err)
			break
		}
		if leftVelocity != 0 || rightVelocity != 0 {
			time.Sleep(time.Second * 3)
		} else {
			time.Sleep(time.Millisecond * 500)
		}
	}
}
