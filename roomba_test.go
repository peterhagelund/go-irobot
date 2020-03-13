package irobot

import (
	"net"
	"testing"
)

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
		err := roomba.Start()
		if err != nil {
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
		err := roomba.Safe()
		if err != nil {
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
		err := roomba.Full()
		if err != nil {
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
		err := roomba.Power()
		if err != nil {
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
		err := roomba.Spot()
		if err != nil {
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
		err := roomba.Clean()
		if err != nil {
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
		err := roomba.Max()
		if err != nil {
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
