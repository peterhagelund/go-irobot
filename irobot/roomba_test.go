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

package irobot

import (
	"net"
	"testing"
	"time"
)

func TestNewNote(t *testing.T) {
	note, err := NewNote(31, 500*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if note.Number != 31 {
		t.Fatal("Number is invalid")
	}
	if note.Duration != 32 {
		t.Fatal("Duration is invalid")
	}
	note, err = NewNote(127, 3000*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if note.Number != 127 {
		t.Fatal("Number is invalid")
	}
	if note.Duration != 192 {
		t.Fatal("Duration is invalid")
	}
	if _, err := NewNote(4, 3000*time.Millisecond); err == nil {
		t.Fatal("invalid number not rejected")
	}
	if _, err := NewNote(31, 4000*time.Millisecond); err == nil {
		t.Fatal("invalid duration not rejected")
	}
}

func TestNewRoomba(t *testing.T) {
	if _, err := NewRoomba(nil); err == nil {
		t.Fatal("nil conn not rejected")
	}
	dummy, conn := net.Pipe()
	defer dummy.Close()
	defer conn.Close()
	roomba, err := NewRoomba(conn)
	if err != nil {
		t.Fatal(err)
	}
	if roomba == nil {
		t.Fatal("nil roomba returned")
	}
}

func TestStart(t *testing.T) {
	dummy, conn := net.Pipe()
	defer dummy.Close()
	defer conn.Close()
	roomba, _ := NewRoomba(conn)
	data := make([]byte, 1)
	go func() {
		if err := roomba.Start(); err != nil {
			t.Error(err)
		}
	}()
	n, err := dummy.Read(data)
	if err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("expected 1 byte, got %d", n)
	}
	if data[0] != 128 {
		t.Fatalf("expected start op-code, got %d", data[0])
	}
}

func TestSetBaudRate(t *testing.T) {
	dummy, conn := net.Pipe()
	defer dummy.Close()
	defer conn.Close()
	roomba, _ := NewRoomba(conn)
	if err := roomba.SetBaudRate(BaudRate57600); err == nil {
		t.Fatal("set baud not rejected on non-port connection")
	}

}

func TestSafe(t *testing.T) {
	dummy, conn := net.Pipe()
	defer dummy.Close()
	defer conn.Close()
	roomba, _ := NewRoomba(conn)
	data := make([]byte, 1)
	go func() {
		if err := roomba.Safe(); err != nil {
			t.Error(err)
		}
	}()
	n, err := dummy.Read(data)
	if err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("expected 1 byte, got %d", n)
	}
	if data[0] != 131 {
		t.Fatalf("expected safe op-code, got %d", data[0])
	}
}

func TestFull(t *testing.T) {
	dummy, conn := net.Pipe()
	defer dummy.Close()
	defer conn.Close()
	roomba, _ := NewRoomba(conn)
	data := make([]byte, 1)
	go func() {
		if err := roomba.Full(); err != nil {
			t.Error(err)
		}
	}()
	n, err := dummy.Read(data)
	if err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("expected 1 byte, got %d", n)
	}
	if data[0] != 132 {
		t.Fatalf("expected full op-code, got %d", data[0])
	}
}

func TestPower(t *testing.T) {
	dummy, conn := net.Pipe()
	defer dummy.Close()
	defer conn.Close()
	roomba, _ := NewRoomba(conn)
	data := make([]byte, 1)
	go func() {
		if err := roomba.Power(); err != nil {
			t.Error(err)
		}
	}()
	n, err := dummy.Read(data)
	if err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("expected 1 byte, got %d", n)
	}
	if data[0] != 133 {
		t.Fatalf("expected power op-code, got %d", data[0])
	}
}

func TestSetMode(t *testing.T) {
	dummy, conn := net.Pipe()
	defer dummy.Close()
	defer conn.Close()
	roomba, _ := NewRoomba(conn)
	if err := roomba.SetMode(ModeUnknown); err == nil {
		t.Fatal("set mode to unknown not rejected")
	}
	data := make([]byte, 1)
	go func() {
		if err := roomba.SetMode(ModeSafe); err != nil {
			t.Error(err)
		}
	}()
	n, err := dummy.Read(data)
	if err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("expected 1 byte, got %d", n)
	}
	if data[0] != 131 {
		t.Fatalf("expected safe op-code, got %d", data[0])
	}
}

func TestSpot(t *testing.T) {
	dummy, conn := net.Pipe()
	defer dummy.Close()
	defer conn.Close()
	roomba, _ := NewRoomba(conn)
	data := make([]byte, 1)
	go func() {
		if err := roomba.Spot(); err != nil {
			t.Error(err)
		}
	}()
	n, err := dummy.Read(data)
	if err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("expected 1 byte, got %d", n)
	}
	if data[0] != 134 {
		t.Fatalf("expected spot op-code, got %d", data[0])
	}
}

func TestClean(t *testing.T) {
	dummy, conn := net.Pipe()
	defer dummy.Close()
	defer conn.Close()
	roomba, _ := NewRoomba(conn)
	data := make([]byte, 1)
	go func() {
		if err := roomba.Clean(); err != nil {
			t.Error(err)
		}
	}()
	n, err := dummy.Read(data)
	if err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("expected 1 byte, got %d", n)
	}
	if data[0] != 135 {
		t.Fatalf("expected clean op-code, got %d", data[0])
	}
}

func TestMax(t *testing.T) {
	dummy, conn := net.Pipe()
	defer dummy.Close()
	defer conn.Close()
	roomba, _ := NewRoomba(conn)
	data := make([]byte, 1)
	go func() {
		if err := roomba.Max(); err != nil {
			t.Error(err)
		}
	}()
	n, err := dummy.Read(data)
	if err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("expected 1 byte, got %d", n)
	}
	if data[0] != 136 {
		t.Fatalf("expected max op-code, got %d", data[0])
	}
}

func TestDrive(t *testing.T) {
	dummy, conn := net.Pipe()
	defer dummy.Close()
	defer conn.Close()
	roomba, _ := NewRoomba(conn)
	if err := roomba.Drive(-700, 0); err == nil {
		t.Fatal("invalid velocity not rejected")
	}
	if err := roomba.Drive(700, 0); err == nil {
		t.Fatal("invalid velocity not rejected")
	}
	if err := roomba.Drive(200, -10000); err == nil {
		t.Fatal("invalid radius not rejected")
	}
	if err := roomba.Drive(200, 10000); err == nil {
		t.Fatal("invalid radius not rejected")
	}
	data := make([]byte, 5)
	go func() {
		if err := roomba.Drive(200, -1000); err != nil {
			t.Error(err)
		}
	}()
	n, err := dummy.Read(data)
	if err != nil {
		t.Fatal(err)
	}
	if n != 5 {
		t.Fatalf("expected 5 bytes, got %d", n)
	}
	if data[0] != 137 {
		t.Fatalf("expected drive op-code, got %d", data[0])
	}
	if data[1] != 0x00 || data[2] != 0xc8 {
		t.Fatalf("expected 0x00, 0xc8 velocity, got 0x%02x, 0x%02x", data[1], data[2])
	}
	if data[3] != 0xfc || data[4] != 0x18 {
		t.Fatalf("expected 0xfc, 0x18 radius, got 0x%02x, 0x%02x", data[3], data[4])
	}
}

func TestMotors(t *testing.T) {
	dummy, conn := net.Pipe()
	defer dummy.Close()
	defer conn.Close()
	roomba, _ := NewRoomba(conn)
	data := make([]byte, 2)
	go func() {
		if err := roomba.Motors(MainBrushInward, SideBrushClockwise, VacuumOn); err != nil {
			t.Error(err)
		}
	}()
	n, err := dummy.Read(data)
	if err != nil {
		t.Fatal(err)
	}
	if n != 2 {
		t.Fatalf("expected 2 bytes, got %d", n)
	}
	if data[0] != 138 {
		t.Fatalf("expected motors op-code, got %d", data[0])
	}
	if data[1] != 0b00001111 {
		t.Fatalf("expected motor bits 0b00001111, got 0b%08b", data[1])
	}
}

func TestLEDs(t *testing.T) {
	dummy, conn := net.Pipe()
	roomba, _ := NewRoomba(conn)
	defer dummy.Close()
	defer conn.Close()
	data := make([]byte, 4)
	go func() {
		if err := roomba.LEDs(50, 100, true, false, true, false); err != nil {
			t.Error(err)
		}
	}()
	n, err := dummy.Read(data)
	if err != nil {
		t.Fatal(err)
	}
	if n != 4 {
		t.Fatalf("expected 4 bytes, got %d", n)
	}
	if data[0] != 139 {
		t.Fatalf("expected LEDs op-code, got %d", data[0])
	}
	if data[1] != 0b00000101 {
		t.Fatalf("expected LED bits 0b00001010, got 0b%08b", data[1])
	}
	if data[2] != 50 {
		t.Fatalf("expected color 50, got %d", data[2])
	}
	if data[3] != 100 {
		t.Fatalf("expected intensity 100, got %d", data[3])
	}
}

func TestSetSong(t *testing.T) {
	dummy, conn := net.Pipe()
	defer dummy.Close()
	defer conn.Close()
	roomba, _ := NewRoomba(conn)
	var notes = make([]*Note, 1)
	notes[0], _ = NewNote(31, 1000*time.Millisecond)
	if err := roomba.SetSong(5, notes); err == nil {
		t.Fatal("invalid song number not rejected")
	}
	notes = make([]*Note, 0)
	if err := roomba.SetSong(1, notes); err == nil {
		t.Fatal("empty notes not rejected")
	}
	notes = make([]*Note, 17)
	for i := 0; i < len(notes); i++ {
		notes[i], _ = NewNote(31, 1000*time.Millisecond)
	}
	if err := roomba.SetSong(1, notes); err == nil {
		t.Fatal("too long notes not rejected")
	}
	notes = make([]*Note, 2)
	notes[0], _ = NewNote(31, 1000*time.Millisecond)
	notes[1], _ = NewNote(32, 1500*time.Millisecond)
	data := make([]byte, 7) // 1 + 1 + 1 + 2 * 2
	go func() {
		if err := roomba.SetSong(1, notes); err != nil {
			t.Error(err)
		}
	}()
	n, err := dummy.Read(data)
	if err != nil {
		t.Fatal(err)
	}
	if n != 7 {
		t.Fatalf("expected 7 bytes, got %d", n)
	}
	if data[0] != 140 {
		t.Fatalf("expected song op-code, got %d", data[0])
	}
	if data[1] != 1 {
		t.Fatalf("expected song number 1, got %d", data[1])
	}
	if data[2] != 2 {
		t.Fatalf("expected song length 2, got %d", data[2])
	}
	if data[3] != 31 {
		t.Fatalf("expected note 31, got %d", data[3])
	}
	if data[4] != 64 {
		t.Fatalf("expected duration 64, got %d", data[3])
	}
	if data[5] != 32 {
		t.Fatalf("expected note 31, got %d", data[5])
	}
	if data[6] != 96 {
		t.Fatalf("expected duration 96, got %d", data[3])
	}
}

func TestPlay(t *testing.T) {
	dummy, conn := net.Pipe()
	defer dummy.Close()
	defer conn.Close()
	roomba, _ := NewRoomba(conn)
	if err := roomba.Play(5); err == nil {
		t.Fatal("invalid song number not not rejected")
	}
	data := make([]byte, 2)
	go func() {
		if err := roomba.Play(2); err != nil {
			t.Error(err)
		}
	}()
	n, err := dummy.Read(data)
	if err != nil {
		t.Fatal(err)
	}
	if n != 2 {
		t.Fatalf("expected 2 bytes, got %d", n)
	}
	if data[0] != 141 {
		t.Fatalf("expected play op-code, got %d", data[0])
	}
	if data[1] != 2 {
		t.Fatalf("expected song number 2, got %d", data[1])
	}
}

func TestSensorsSimple(t *testing.T) {
	dummy, conn := net.Pipe()
	defer dummy.Close()
	defer conn.Close()
	roomba, _ := NewRoomba(conn)
	if _, err := roomba.Sensors(111); err == nil {
		t.Fatal("invalid packet id not rejected")
	}
	c := make(chan int)
	go func() {
		data := make([]byte, 2)
		data[0] = 0x3e
		data[1] = 0x80
		if _, err := dummy.Write(data); err != nil {
			t.Error(err)
		}
		c <- 1
	}()
	data := make([]byte, 2)
	go func() {
		packet, err := roomba.Sensors(22)
		if err != nil {
			t.Error(err)
		}
		packet22 := packet.(*Packet22)
		if packet22.Voltage != 16000 {
			t.Errorf("expected voltage 16000, got 0x%02x", packet22.Voltage)
		}
		c <- 2
	}()
	n, err := dummy.Read(data)
	if err != nil {
		t.Fatal(err)
	}
	if n != 2 {
		t.Fatalf("expected 2 bytes, got %d", n)
	}
	if data[0] != 142 {
		t.Fatalf("expected sensors op-code, got %d", data[0])
	}
	if data[1] != 22 {
		t.Fatalf("expected packet id 22, got 0x%02x", data[1])
	}
	<-c
	<-c
}

func TestSensorsGroup(t *testing.T) {
	dummy, conn := net.Pipe()
	defer dummy.Close()
	defer conn.Close()
	roomba, _ := NewRoomba(conn)
	c := make(chan int)
	go func() {
		data := []byte{
			0x40,       // Packet17
			0b00000100, // Packet18
			0x30,       // Packet19
			0x39,       // -
			0xfc,       // Packet20
			0x19,       // -
		}
		if _, err := dummy.Write(data); err != nil {
			t.Error(err)
		}
		c <- 1
	}()
	data := make([]byte, 2)
	go func() {
		packet, err := roomba.Sensors(2)
		if err != nil {
			t.Error(err)
		}
		packet2 := packet.(*Packet2)
		// Validate Packet17
		if packet2.Packet17.InfraredCharacterOmni != 0x40 {
			t.Error("InfraredCharacterOmni has wrong value")
		}
		// Validate Packet18
		if packet2.Packet18.Clean != false {
			t.Error("Clean has wrong value")
		}
		if packet2.Packet18.Spot != false {
			t.Error("Spot has wrong value")
		}
		if packet2.Packet18.Dock != true {
			t.Error("Dock has wrong value")
		}
		if packet2.Packet18.Minute != false {
			t.Error("Minute has wrong value")
		}
		if packet2.Packet18.Hour != false {
			t.Error("Hour has wrong value")
		}
		if packet2.Packet18.Day != false {
			t.Error("Day has wrong value")
		}
		if packet2.Packet18.Schedule != false {
			t.Error("Schedule has wrong value")
		}
		if packet2.Packet18.Clock != false {
			t.Error("Clock has wrong value")
		}
		// Validate Packet19
		if packet2.Packet19.Distance != 12345 {
			t.Error("Distance has wrong value")
		}
		// Validate Packet20
		if packet2.Packet20.Angle != -999 {
			t.Error("Angle has wrong value")
		}
		c <- 2
	}()
	n, err := dummy.Read(data)
	if err != nil {
		t.Fatal(err)
	}
	if n != 2 {
		t.Fatalf("expected 2 bytes, got %d", n)
	}
	if data[0] != 142 {
		t.Fatalf("expected sensors op-code, got %d", data[0])
	}
	if data[1] != 2 {
		t.Fatalf("expected packet id 2, got 0x%02x", data[1])
	}
	<-c
	<-c
}

func TestMotorsPWM(t *testing.T) {
	dummy, conn := net.Pipe()
	defer dummy.Close()
	defer conn.Close()
	roomba, _ := NewRoomba(conn)
	if err := roomba.MotorsPWM(-128, 0, 0); err == nil {
		t.Fatal("invalid main brush PWM not rejected")
	}
	if err := roomba.MotorsPWM(0, -128, 0); err == nil {
		t.Fatal("invalid side brush PWM not rejected")
	}
	if err := roomba.MotorsPWM(0, 0, 128); err == nil {
		t.Fatal("invalid vacuum PWM not rejected")
	}
	data := make([]byte, 4)
	go func() {
		if err := roomba.MotorsPWM(-50, 50, 100); err != nil {
			t.Error(err)
		}
	}()
	n, err := dummy.Read(data)
	if err != nil {
		t.Fatal(err)
	}
	if n != 4 {
		t.Fatalf("expected 4 bytes, got %d", n)
	}
	if data[0] != 144 {
		t.Fatalf("expected motors PWN op-code, got %d", data[0])
	}
	if data[1] != 0xce {
		t.Fatalf("expected 0xce main brush PWM, got 0x%02x", data[1])
	}
	if data[2] != 0x32 {
		t.Fatalf("expected 0x32 side brush PWM, got 0x%02x", data[2])
	}
	if data[3] != 0x64 {
		t.Fatalf("expected 0x64 vacuum PWM, got 0x%02x", data[3])
	}
}

func TestDriveDirect(t *testing.T) {
	dummy, conn := net.Pipe()
	defer dummy.Close()
	defer conn.Close()
	roomba, _ := NewRoomba(conn)
	if err := roomba.DriveDirect(-600, 200); err == nil {
		t.Fatal("invalid left velocity not rejected")
	}
	if err := roomba.DriveDirect(600, 200); err == nil {
		t.Fatal("invalid left velocity not rejected")
	}
	if err := roomba.DriveDirect(200, -600); err == nil {
		t.Fatal("invalid right velocity not rejected")
	}
	if err := roomba.DriveDirect(200, 600); err == nil {
		t.Fatal("invalid right velocity not rejected")
	}
	data := make([]byte, 5)
	go func() {
		if err := roomba.DriveDirect(-250, 250); err != nil {
			t.Error(err)
		}
	}()
	n, err := dummy.Read(data)
	if err != nil {
		t.Fatal(err)
	}
	if n != 5 {
		t.Fatalf("expected 5 bytes, got %d", n)
	}
	if data[0] != 145 {
		t.Fatalf("expected drive direct op-code, got %d", data[0])
	}
	if data[1] != 0x00 || data[2] != 0xfa {
		t.Fatalf("expected 0x00, 0xfa right velocity, got 0x%02x, 0x%02x", data[1], data[2])
	}
	if data[3] != 0xff || data[4] != 0x06 {
		t.Fatalf("expected 0xff, 0x06 left velocity, got 0x%02x, 0x%02x", data[3], data[4])
	}
}

func TestDrivePWM(t *testing.T) {
	dummy, conn := net.Pipe()
	defer dummy.Close()
	defer conn.Close()
	roomba, _ := NewRoomba(conn)
	if err := roomba.DrivePWM(-300, 200); err == nil {
		t.Fatal("invalid left wheel PWM not rejected")
	}
	if err := roomba.DrivePWM(300, 200); err == nil {
		t.Fatal("invalid left wheel PWM not rejected")
	}
	if err := roomba.DrivePWM(200, -300); err == nil {
		t.Fatal("invalid right wheel PWM not rejected")
	}
	if err := roomba.DrivePWM(200, 300); err == nil {
		t.Fatal("invalid right wheel PWM not rejected")
	}
	data := make([]byte, 5)
	go func() {
		if err := roomba.DrivePWM(200, -200); err != nil {
			t.Error(err)
		}
	}()
	n, err := dummy.Read(data)
	if err != nil {
		t.Fatal(err)
	}
	if n != 5 {
		t.Fatalf("expected 5 bytes, got %d", n)
	}
	if data[0] != 146 {
		t.Fatalf("expected drive PWM op-code, got %d", data[0])
	}
	if data[1] != 0xff || data[2] != 0x38 {
		t.Fatalf("expected 0xff, 0x38 right wheel PWM, got 0x%02x, 0x%02x", data[1], data[2])
	}
	if data[3] != 0x00 || data[4] != 0xc8 {
		t.Fatalf("expected 0x00, 0xc8 left wheel PWM, got 0x%02x, 0x%02x", data[3], data[4])
	}
}
