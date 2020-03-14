package irobot

import "testing"

func TestExtract19(t *testing.T) {
	data := make([]byte, 2)
	packet := makePacket19().(*Packet19)
	data[0] = 0x00
	data[1] = 0x00
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.Distance != 0 {
		t.Errorf("Distance has wrong value")
	}
	data[0] = 0x01
	data[1] = 0xf4
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.Distance != 500 {
		t.Errorf("Distance has wrong value")
	}
	data[0] = 0xfe
	data[1] = 0x0c
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.Distance != -500 {
		t.Errorf("Distance has wrong value")
	}
}
