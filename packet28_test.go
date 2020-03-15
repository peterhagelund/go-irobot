package irobot

import "testing"

func TestExtract28(t *testing.T) {
	packet := makePacket28().(*Packet28)
	data := make([]byte, 2)
	data[0] = 0x00
	data[1] = 0x00
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.CliffLeftSignal != 0 {
		t.Errorf("CliffLeftSignal has wrong value")
	}
	data[0] = 0x0f
	data[1] = 0xff
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.CliffLeftSignal != 4095 {
		t.Errorf("CliffLeftSignal has wrong value")
	}
	data[0] = 0x10
	data[1] = 0x00
	if err := packet.Extract(data, 0); err == nil {
		t.Error("invalid cliff left signal not rejected")
	}
}
