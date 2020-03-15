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
		t.Error("Number is invalid")
	}
	if note.Duration != 32 {
		t.Error("Duration is invalid")
	}
	note, err = NewNote(127, 3000*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if note.Number != 127 {
		t.Error("Number is invalid")
	}
	if note.Duration != 192 {
		t.Error("Duration is invalid")
	}
	_, err = NewNote(31, 4000*time.Millisecond)
	if err == nil {
		t.Error("invalid duration not rejected")
	}
}

func TestNewRoomba(t *testing.T) {
	_, conn := net.Pipe()
	roomba, err := NewRoomba(conn)
	if err != nil {
		t.Error(err)
	}
	if roomba == nil {
		t.Error("nil roomba returned")
	}
}

func TestStart(t *testing.T) {
	dummy, conn := net.Pipe()
	roomba, _ := NewRoomba(conn)
	data := make([]byte, 1)
	go func() {
		if err := roomba.Start(); err != nil {
			t.Error(err)
		}
	}()
	n, err := dummy.Read(data)
	if err != nil {
		t.Error(err)
	}
	if n != 1 {
		t.Errorf("expected 1 byte, got %d", n)
	}
	if data[0] != 128 {
		t.Errorf("expected start op-code, got %d", data[0])
	}
}

func TestSafe(t *testing.T) {
	dummy, conn := net.Pipe()
	roomba, _ := NewRoomba(conn)
	data := make([]byte, 1)
	go func() {
		if err := roomba.Safe(); err != nil {
			t.Error(err)
		}
	}()
	n, err := dummy.Read(data)
	if err != nil {
		t.Error(err)
	}
	if n != 1 {
		t.Errorf("expected 1 byte, got %d", n)
	}
	if data[0] != 131 {
		t.Errorf("expected safe op-code, got %d", data[0])
	}
}

func TestFull(t *testing.T) {
	dummy, conn := net.Pipe()
	roomba, _ := NewRoomba(conn)
	data := make([]byte, 1)
	go func() {
		if err := roomba.Full(); err != nil {
			t.Error(err)
		}
	}()
	n, err := dummy.Read(data)
	if err != nil {
		t.Error(err)
	}
	if n != 1 {
		t.Errorf("expected 1 byte, got %d", n)
	}
	if data[0] != 132 {
		t.Errorf("expected full op-code, got %d", data[0])
	}
}

func TestPower(t *testing.T) {
	dummy, conn := net.Pipe()
	roomba, _ := NewRoomba(conn)
	data := make([]byte, 1)
	go func() {
		if err := roomba.Power(); err != nil {
			t.Error(err)
		}
	}()
	n, err := dummy.Read(data)
	if err != nil {
		t.Error(err)
	}
	if n != 1 {
		t.Errorf("expected 1 byte, got %d", n)
	}
	if data[0] != 133 {
		t.Errorf("expected power op-code, got %d", data[0])
	}
}

func TestSpot(t *testing.T) {
	dummy, conn := net.Pipe()
	roomba, _ := NewRoomba(conn)
	data := make([]byte, 1)
	go func() {
		if err := roomba.Spot(); err != nil {
			t.Error(err)
		}
	}()
	n, err := dummy.Read(data)
	if err != nil {
		t.Error(err)
	}
	if n != 1 {
		t.Errorf("expected 1 byte, got %d", n)
	}
	if data[0] != 134 {
		t.Errorf("expected spot op-code, got %d", data[0])
	}
}

func TestClean(t *testing.T) {
	dummy, conn := net.Pipe()
	roomba, _ := NewRoomba(conn)
	data := make([]byte, 1)
	go func() {
		if err := roomba.Clean(); err != nil {
			t.Error(err)
		}
	}()
	n, err := dummy.Read(data)
	if err != nil {
		t.Error(err)
	}
	if n != 1 {
		t.Errorf("expected 1 byte, got %d", n)
	}
	if data[0] != 135 {
		t.Errorf("expected clean op-code, got %d", data[0])
	}
}

func TestMax(t *testing.T) {
	dummy, conn := net.Pipe()
	roomba, _ := NewRoomba(conn)
	data := make([]byte, 1)
	go func() {
		if err := roomba.Max(); err != nil {
			t.Error(err)
		}
	}()
	n, err := dummy.Read(data)
	if err != nil {
		t.Error(err)
	}
	if n != 1 {
		t.Errorf("expected 1 byte, got %d", n)
	}
	if data[0] != 136 {
		t.Errorf("expected max op-code, got %d", data[0])
	}
}

func TestDrive(t *testing.T) {
	dummy, conn := net.Pipe()
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
		t.Error(err)
	}
	if n != 5 {
		t.Errorf("expected 1 byte, got %d", n)
	}
	if data[0] != 137 {
		t.Errorf("expected drive op-code, got %d", data[0])
	}
	if data[1] != 0x00 || data[2] != 0xc8 {
		t.Errorf("expected 0x00, 0xc8 velocity, got 0x%02x, 0x%02x", data[1], data[2])
	}
	if data[3] != 0xfc || data[4] != 0x18 {
		t.Errorf("expected 0xfc, 0x18 radius, got 0x%02x, 0x%02x", data[3], data[4])
	}
}

func TestMotors(t *testing.T) {
	dummy, conn := net.Pipe()
	roomba, _ := NewRoomba(conn)
	data := make([]byte, 2)
	go func() {
		if err := roomba.Motors(MainBrushInward, SideBrushClockwise, VacuumOn); err != nil {
			t.Error(err)
		}
	}()
	n, err := dummy.Read(data)
	if err != nil {
		t.Error(err)
	}
	if n != 2 {
		t.Errorf("expected 1 byte, got %d", n)
	}
	if data[0] != 138 {
		t.Errorf("expected motors op-code, got %d", data[0])
	}
	if data[1] != 0b00001111 {
		t.Errorf("expected motor bits 0b00001111, got 0b%08b", data[1])
	}
}

func TestDriveDirect(t *testing.T) {
	dummy, conn := net.Pipe()
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
		t.Error(err)
	}
	if n != 5 {
		t.Errorf("expected 1 byte, got %d", n)
	}
	if data[0] != 145 {
		t.Errorf("expected drive op-code, got %d", data[0])
	}
	if data[1] != 0x00 || data[2] != 0xfa {
		t.Errorf("expected 0x00, 0xfa right velocity, got 0x%02x, 0x%02x", data[1], data[2])
	}
	if data[3] != 0xff || data[4] != 0x06 {
		t.Errorf("expected 0xff, 0x06 left velocity, got 0x%02x, 0x%02x", data[3], data[4])
	}
}
