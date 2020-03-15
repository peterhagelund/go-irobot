package irobot

import "testing"

func TestExtract22(t *testing.T) {
	packet := makePacket22().(*Packet22)
	data := make([]byte, 2)
	data[0] = 0x00
	data[1] = 0x00
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.Voltage != 0 {
		t.Errorf("Voltage has wrong value")
	}
	data[0] = 0x42
	data[1] = 0x68
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.Voltage != 17000 {
		t.Errorf("Voltage has wrong value")
	}
}
