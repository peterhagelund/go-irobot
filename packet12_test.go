package irobot

import "testing"

func TestExtract12(t *testing.T) {
	data := make([]byte, 1)
	packet := makePacket12().(*Packet12)
	data[0] = 0b00000000
	err := packet.Extract(data, 0)
	if err != nil {
		t.Error(err)
	}
	if packet.CliffRight != false {
		t.Errorf("CliffRight has wrong value")
	}
	data[0] = 0b00000001
	err = packet.Extract(data, 0)
	if err != nil {
		t.Error(err)
	}
	if packet.CliffRight != true {
		t.Errorf("CliffRight has wrong value")
	}
}
