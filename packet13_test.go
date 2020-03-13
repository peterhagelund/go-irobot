package irobot

import "testing"

func TestExtract13(t *testing.T) {
	data := make([]byte, 1)
	packet := makePacket13().(*Packet13)
	data[0] = 0b00000000
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.VirtualWall != false {
		t.Errorf("VirtualWall has wrong value")
	}
	data[0] = 0b00000001
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.VirtualWall != true {
		t.Errorf("VirtualWall has wrong value")
	}
}
