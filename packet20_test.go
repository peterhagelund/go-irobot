package irobot

import "testing"

func TestExtract20(t *testing.T) {
	packet := makePacket20().(*Packet20)
	data := make([]byte, 2)
	data[0] = 0x00
	data[1] = 0x00
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.Angle != 0 {
		t.Errorf("Angle has wrong value")
	}
	data[0] = 0x00
	data[1] = 0xc8
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.Angle != 200 {
		t.Errorf("Angle has wrong value")
	}
	data[0] = 0xff
	data[1] = 0x38
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.Angle != -200 {
		t.Errorf("Angle has wrong value")
	}
}
