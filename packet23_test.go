package irobot

import "testing"

func TestExtract23(t *testing.T) {
	data := make([]byte, 2)
	packet := makePacket23().(*Packet23)
	data[0] = 0x00
	data[1] = 0x00
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.Current != 0 {
		t.Errorf("Current has wrong value")
	}
	data[0] = 0x0f
	data[1] = 0xa0
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.Current != 4000 {
		t.Errorf("Current has wrong value")
	}
	data[0] = 0xf0
	data[1] = 0x60
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.Current != -4000 {
		t.Errorf("Current has wrong value")
	}
}
