# iRobot

Library for controlling [iRobot](https://www.irobot.com) Roomba robots.

## Copyright and License

Copyright (c) 2020 Peter Hagelund

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
years (my first one was a _Roomba Red_) and have written libraries similar to this one in Java, C, C++, Python, Go, and likely other
languages simply because I enjoy doing so.

## Installing

```bash
go get -u github.com/peterhagelund/go-irobot
```

## Using
```go
package main

import (
    "github.com/peterhagelund/go-irobot"
    "github.com/peterhagelund/go-serial"
)

func main() {
	conn, err := serial.Dial("/dev/tty.usbserial-AC01A7BB", serial.BaudRate115200, serial.ParityNone, serial.DataBits8, serial.StopBits1)
	if err != nil {
		panic(err)
	}
    defer conn.Close()
    roomba, err := irobot.NewRoomba(conn)
	if err != nil {
		panic(err)
    }
    // TODO - more example code.
	port.Write([]byte("Hello World via serial\n"))
}
```

_Please note:_ If your code fails at any time after instructing the `Roomba` to drive, it will _continue to drive_ until some piece of code stops it,
so (much more) error handling (see below) is needed than what is described above.

## Error Handling

Under the covers the `Roomba` communicates with a real-world, physical vacuum cleaner, so anything can go wrong at any time.

All public functions return an `error` and in case of an error.
