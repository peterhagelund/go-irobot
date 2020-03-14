package irobot

import "testing"

func TestExtract29(t *testing.T) {
	data := make([]byte, 2)
	packet := makePacket29().(*Packet29)
	data[0] = 0x00
	data[1] = 0x00
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.CliffFrontLeftSignal != 0 {
		t.Errorf("CliffFrontLeftSignal has wrong value")
	}
	data[0] = 0x0f
	data[1] = 0xff
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.CliffFrontLeftSignal != 4095 {
		t.Errorf("CliffFrontLeftSignal has wrong value")
	}
	data[0] = 0x10
	data[1] = 0x00
	if err := packet.Extract(data, 0); err == nil {
		t.Error("invalid cliff front left signal not rejected")
	}
}
