package irobot

import "testing"

func TestExtract26(t *testing.T) {
	data := make([]byte, 2)
	packet := makePacket26().(*Packet26)
	data[0] = 0x00
	data[1] = 0x00
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.BatteryCapacity != 0 {
		t.Errorf("BatteryCapacity has wrong value")
	}
	data[0] = 0x13
	data[1] = 0x88
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.BatteryCapacity != 5000 {
		t.Errorf("BatteryCapacity has wrong value")
	}
}
