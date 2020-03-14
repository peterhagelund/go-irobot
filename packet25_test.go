package irobot

import "testing"

func TestExtract25(t *testing.T) {
	data := make([]byte, 2)
	packet := makePacket25().(*Packet25)
	data[0] = 0x00
	data[1] = 0x00
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.BatteryCharge != 0 {
		t.Errorf("BatteryCharge has wrong value")
	}
	data[0] = 0x13
	data[1] = 0x88
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.BatteryCharge != 5000 {
		t.Errorf("BatteryCharge has wrong value")
	}
}
