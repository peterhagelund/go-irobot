package irobot

import "testing"

func TestExtract32(t *testing.T) {
	packet := makePacket32().(*Packet32)
	data := make([]byte, 1)
	data[0] = 0
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.UnusedByte != 0 {
		t.Errorf("UnusedByte has wrong value")
	}
	data[0] = 42
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.UnusedByte != 42 {
		t.Errorf("UnusedByte has wrong value")
	}
}
