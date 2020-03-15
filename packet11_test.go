package irobot

import "testing"

func TestExtract11(t *testing.T) {
	packet := makePacket11().(*Packet11)
	data := make([]byte, 1)
	data[0] = 0b00000000
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.CliffFrontRight != false {
		t.Errorf("CliffFrontRight has wrong value")
	}
	data[0] = 0b00000001
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.CliffFrontRight != true {
		t.Errorf("CliffFrontRight has wrong value")
	}
}
