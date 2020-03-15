package irobot

import "testing"

func TestExtract34(t *testing.T) {
	data := make([]byte, 1)
	packet := makePacket34().(*Packet34)
	data[0] = 0b00000000
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.InternalCharger != false {
		t.Errorf("InternalCharger has wrong value")
	}
	if packet.HomeBase != false {
		t.Errorf("HomeBase has wrong value")
	}
	data[0] = 0b00000011
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.InternalCharger != true {
		t.Errorf("InternalCharger has wrong value")
	}
	if packet.HomeBase != true {
		t.Errorf("HomeBase has wrong value")
	}
	data[0] = 0b00000010
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.InternalCharger != false {
		t.Errorf("InternalCharger has wrong value")
	}
	if packet.HomeBase != true {
		t.Errorf("HomeBase has wrong value")
	}
}
