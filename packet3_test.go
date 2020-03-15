package irobot

import "testing"

func TestExtract3(t *testing.T) {
	packet := makePacket3().(*Packet3)
	data := make([]byte, 10)
	data[0] = 0x03 // Packet21
	data[1] = 0x40 // Packet22
	data[2] = 0x9f // -
	data[3] = 0xfb // Packet23
	data[4] = 0x2e // -
	data[5] = 19   // Packet24
	data[6] = 0x13 // Packet25
	data[7] = 0x88 // -
	data[8] = 0x27 // Packet26
	data[9] = 0x10 // -
	err := packet.Extract(data, 0)
	if err != nil {
		t.Error(err)
	}
	// Validate Packet21
	if packet.Packet21.ChargingState != ChargingStateTrickleCharging {
		t.Error("ChargingState has wrong value")
	}
	// Validate Packet22
	if packet.Packet22.Voltage != 16543 {
		t.Error("Voltage has wrong value")
	}
	// Validate Packet23
	if packet.Packet23.Current != -1234 {
		t.Error("Current has wrong value")
	}
	// Validate Packet24
	if packet.Packet24.Temperature != 19 {
		t.Error("Temperature has wrong value")
	}
	// Validate Packet25
	if packet.Packet25.BatteryCharge != 5000 {
		t.Error("BatteryCharge has wrong value")
	}
	// Validate Packet26
	if packet.Packet26.BatteryCapacity != 10000 {
		t.Error("BatteryCapacity has wrong value")
	}
}
