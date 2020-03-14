package irobot

import "testing"

func TestExtract27(t *testing.T) {
	data := make([]byte, 2)
	packet := makePacket27().(*Packet27)
	data[0] = 0x00
	data[1] = 0x00
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.WallSignal != 0 {
		t.Errorf("WallSignal has wrong value")
	}
	data[0] = 0x03
	data[1] = 0xff
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.WallSignal != 1023 {
		t.Errorf("WallSignal has wrong value")
	}
	data[0] = 0x04
	data[1] = 0x00
	if err := packet.Extract(data, 0); err == nil {
		t.Error("invalid wall signal not rejected")
	}
}
