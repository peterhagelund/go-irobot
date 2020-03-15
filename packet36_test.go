package irobot

import "testing"

func TestExtract36(t *testing.T) {
	packet := makePacket36().(*Packet36)
	data := make([]byte, 1)
	data[0] = 0
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.Song != 0 {
		t.Errorf("Song has wrong value")
	}
	data[0] = 15
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.Song != 15 {
		t.Errorf("Song has wrong value")
	}
	data[0] = 16
	if err := packet.Extract(data, 0); err == nil {
		t.Error("invalid song not rejected")
	}
}
