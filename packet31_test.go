package irobot

import "testing"

func TestExtract31(t *testing.T) {
	data := make([]byte, 2)
	packet := makePacket31().(*Packet31)
	data[0] = 0x00
	data[1] = 0x00
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.CliffRightSignal != 0 {
		t.Errorf("CliffRightSignal has wrong value")
	}
	data[0] = 0x0f
	data[1] = 0xff
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.CliffRightSignal != 4095 {
		t.Errorf("CliffRightSignal has wrong value")
	}
	data[0] = 0x10
	data[1] = 0x00
	if err := packet.Extract(data, 0); err == nil {
		t.Error("invalid cliff right signal not rejected")
	}
}
