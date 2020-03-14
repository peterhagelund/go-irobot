package irobot

import "testing"

func TestExtract30(t *testing.T) {
	data := make([]byte, 2)
	packet := makePacket30().(*Packet30)
	data[0] = 0x00
	data[1] = 0x00
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.CliffFrontRightSignal != 0 {
		t.Errorf("CliffFrontRightSignal has wrong value")
	}
	data[0] = 0x0f
	data[1] = 0xff
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.CliffFrontRightSignal != 4095 {
		t.Errorf("CliffFrontRightSignal has wrong value")
	}
	data[0] = 0x10
	data[1] = 0x00
	if err := packet.Extract(data, 0); err == nil {
		t.Error("invalid cliff front right signal not rejected")
	}
}
