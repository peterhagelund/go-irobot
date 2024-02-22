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
	defer port.Close()
	fmt.Println("Connected!")
	port.SetReadTimeout(time.Second * 1)
	roomba, err := irobot.NewRoomba(port)
	if err != nil {
		log.Fatal(err)
	}
	if err = roomba.Start(); err != nil {
		log.Fatal(err)
	}
	packet, err := roomba.Sensors(22)
	if err != nil {
		log.Fatal(err)
	}
	packet22, ok := packet.(*irobot.Packet22)
	if !ok {
		log.Fatal("wrong packet")
	}
	for i := range 5 {
		fmt.Printf("voltage %d = %d\n", i, packet22.Voltage)
		time.Sleep(time.Second * 2)
		err := roomba.UpdateSensors(packet22)
		if err != nil {
			fmt.Printf("err = %v\n", err)
		}
	}
	roomba.Power()
}
