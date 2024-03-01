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
	fmt.Println("Started.")
	packet22, _ := irobot.NewPacket[*irobot.Packet22]()
	packet23, _ := irobot.NewPacket[*irobot.Packet23]()
	packet24, _ := irobot.NewPacket[*irobot.Packet24]()
	packet25, _ := irobot.NewPacket[*irobot.Packet25]()
	packet26, _ := irobot.NewPacket[*irobot.Packet26]()
	packets := []irobot.Packet{packet22, packet23, packet24, packet25, packet26}
	for i := range 5 {
		fmt.Println("Querying power information...")
		err := roomba.UpdateList(packets)
		if err != nil {
			fmt.Printf("Error updating packets: %v\n", err)
		} else {
			fmt.Printf("Voltage %d .......: %d\n", i, packet22.Voltage)
			fmt.Printf("Current %d .......: %d\n", i, packet23.Current)
			fmt.Printf("Temperature %d ...: %d\n", i, packet24.Temperature)
			fmt.Printf("Charge %d ........: %d\n", i, packet25.BatteryCharge)
			fmt.Printf("Capacity %d ......: %d\n", i, packet26.BatteryCapacity)
		}
		fmt.Println("Sleeping...")
		time.Sleep(time.Second * 1)
	}
	roomba.Power()
	fmt.Println("Done.")
}
