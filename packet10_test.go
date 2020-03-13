package irobot

import "testing"

func TestExtract10(t *testing.T) {
	data := make([]byte, 1)
	packet := makePacket10().(*Packet10)
	data[0] = 0b00000000
	err := packet.Extract(data[:], 0)
	if err != nil {
		t.Error(err)
	}
	if packet.CliffFrontLeft != false {
		t.Errorf("CliffFrontLeft has wrong value")
	}
	data[0] = 0b00000001
	err = packet.Extract(data[:], 0)
	if err != nil {
		t.Error(err)
	}
	if packet.CliffFrontLeft != true {
		t.Errorf("CliffFrontLeft has wrong value")
	}
}
