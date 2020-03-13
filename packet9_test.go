package irobot

import "testing"

func TestExtract9(t *testing.T) {
	data := make([]byte, 1)
	packet := makePacket9().(*Packet9)
	data[0] = 0b00000000
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.CliffLeft != false {
		t.Errorf("CliffLeft has wrong value")
	}
	data[0] = 0b00000001
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.CliffLeft != true {
		t.Errorf("CliffLeft has wrong value")
	}
}
