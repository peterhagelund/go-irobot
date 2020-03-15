package irobot

import "testing"

func TestExtract12(t *testing.T) {
	packet := makePacket12().(*Packet12)
	data := make([]byte, 1)
	data[0] = 0b00000000
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.CliffRight != false {
		t.Errorf("CliffRight has wrong value")
	}
	data[0] = 0b00000001
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.CliffRight != true {
		t.Errorf("CliffRight has wrong value")
	}
}
