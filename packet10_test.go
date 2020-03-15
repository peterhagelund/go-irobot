package irobot

import "testing"

func TestExtract10(t *testing.T) {
	packet := makePacket10().(*Packet10)
	data := make([]byte, 1)
	data[0] = 0b00000000
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.CliffFrontLeft != false {
		t.Errorf("CliffFrontLeft has wrong value")
	}
	data[0] = 0b00000001
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.CliffFrontLeft != true {
		t.Errorf("CliffFrontLeft has wrong value")
	}
}
