package irobot

import "testing"

func TestExtract33(t *testing.T) {
	data := make([]byte, 2)
	packet := makePacket33().(*Packet33)
	data[0] = 0x00
	data[1] = 0x00
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.UnusedWord != 0 {
		t.Errorf("UnusedWord has wrong value")
	}
	data[0] = 0x12
	data[1] = 0x34
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.UnusedWord != 4660 {
		t.Errorf("UnusedWord has wrong value")
	}
}
