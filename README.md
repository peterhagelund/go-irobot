# iRobot

Library for controlling [iRobot](https://www.irobot.com) Roomba robots.

## Copyright and License

Copyright (c) 2020-2024 Peter Hagelund

This software is licensed under the [MIT License](https://en.wikipedia.org/wiki/MIT_License)

See `LICENSE.txt`

_Please note:_ Pay particular attention to the bottom half of the MIT license that talks about the software being provided "AS IS"
and that there is no warranty of any kind. If you choose to utilize this software to control (or at least attempt to control) a physical
device in the real world, you run the risk of damaging the vaccum cleaner and/or your dwelling and/or any furniture or other
belongings you may keep there and/or any carbon-based lifeforms that might inhabit your space. This is **especially** true if you
choose to place the `Roomba` OI in `full` mode - in which case all of the failsafes the talented iRobot engineers have put in place
are *completely disabled*. In `full` mode you run the risk of not catching an overcurrent condition in the motors, having the vacuum
cleaner fall down a staircase, and possibly hit and ingest your pet(s) or a child. Read the [OI specification](https://www.irobot.lv/uploaded_files/File/iRobot_Roomba_500_Open_Interface_Spec.pdf)
very carefully and _test your software to destruction_ **before** letting this thing loose in your house. _Ye be warned!_

_Please note:_ I am in no way affiliated with iRobot or any of its subsidiaries or suppliers. I have used their products for several
years (my first one was a _Roomba Red_) and have written libraries similar to this one in Java, C, C++, Python, Swift, and likely other languages simply because I enjoy doing so.

## Installing

```bash
go get -u github.com/peterhagelund/go-irobot
```

## Using
```go
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
```

_Please note:_ If your code fails at any time after instructing the `Roomba` to drive, it will _continue to drive_ until some piece of code stops it,
so (much more) error handling (see below) is needed than what is described above.

## Error Handling

Under the covers the `Roomba` communicates with a real-world, physical vacuum cleaner, so anything can go wrong at any time.

All public functions return an `error` and in case of an error.
