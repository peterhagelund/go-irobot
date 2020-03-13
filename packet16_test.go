package irobot

import "testing"

func TestExtract16(t *testing.T) {
	data := make([]byte, 1)
	packet := makePacket16().(*Packet16)
	data[0] = 0
	err := packet.Extract(data, 0)
	if err != nil {
		t.Error(err)
	}
	if packet.UnusedByte != 0 {
		t.Errorf("UnusedByte has wrong value")
	}
	data[0] = 42
	err = packet.Extract(data, 0)
	if err != nil {
		t.Error(err)
	}
	if packet.UnusedByte != 42 {
		t.Errorf("UnusedByte has wrong value")
	}
}