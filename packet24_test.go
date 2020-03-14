package irobot

import "testing"

func TestExtract24(t *testing.T) {
	data := make([]byte, 1)
	packet := makePacket24().(*Packet24)
	data[0] = 0
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.Temperature != 0 {
		t.Errorf("Temperature has wrong value")
	}
	data[0] = 20
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.Temperature != 20 {
		t.Errorf("Temperature has wrong value")
	}
}
