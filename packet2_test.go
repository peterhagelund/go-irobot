package irobot

import "testing"

func TestExtract2(t *testing.T) {
	packet := makePacket2().(*Packet2)
	data := make([]byte, 6)
	data[0] = 0x40       // Packet17
	data[1] = 0b00000100 // Packet18
	data[2] = 0x30       // Packet19
	data[3] = 0x39       // -
	data[4] = 0xfc       // Packet20
	data[5] = 0x19       // -
	err := packet.Extract(data, 0)
	if err != nil {
		t.Error(err)
	}
	// Validate Packet17
	if packet.Packet17.InfraredCharacterOmni != 0x40 {
		t.Error("InfraredCharacterOmni has wrong value")
	}
	// Validate Packet18
	if packet.Packet18.Clean != false {
		t.Error("Clean has wrong value")
	}
	if packet.Packet18.Spot != false {
		t.Error("Spot has wrong value")
	}
	if packet.Packet18.Dock != true {
		t.Error("Dock has wrong value")
	}
	if packet.Packet18.Minute != false {
		t.Error("Minute has wrong value")
	}
	if packet.Packet18.Hour != false {
		t.Error("Hour has wrong value")
	}
	if packet.Packet18.Day != false {
		t.Error("Day has wrong value")
	}
	if packet.Packet18.Schedule != false {
		t.Error("Schedule has wrong value")
	}
	if packet.Packet18.Clock != false {
		t.Error("Clock has wrong value")
	}
	// Validate Packet19
	if packet.Packet19.Distance != 12345 {
		t.Error("Distance has wrong value")
	}
	// Validate Packet20
	if packet.Packet20.Angle != -999 {
		t.Error("Angle has wrong value")
	}
}
