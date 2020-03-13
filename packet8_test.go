package irobot

import "testing"

func TestExtract8(t *testing.T) {
	data := make([]byte, 1)
	packet := makePacket8().(*Packet8)
	data[0] = 0b00000000
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.Wall != false {
		t.Errorf("Wall has wrong value")
	}
	data[0] = 0b00000001
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.Wall != true {
		t.Errorf("Wall has wrong value")
	}
}
