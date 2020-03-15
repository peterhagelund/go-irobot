package irobot

import "testing"

func TestExtract35(t *testing.T) {
	packet := makePacket35().(*Packet35)
	data := make([]byte, 1)
	data[0] = 0
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.Mode != ModeOff {
		t.Errorf("Mode has wrong value")
	}
	data[0] = 2
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.Mode != ModeSafe {
		t.Errorf("Mode has wrong value")
	}
	data[0] = 42
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.Mode != ModeUnknown {
		t.Errorf("Mode has wrong value")
	}
}
