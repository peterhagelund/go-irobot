package irobot

import "testing"

func TestExtract18(t *testing.T) {
	packet := makePacket18().(*Packet18)
	data := make([]byte, 1)
	data[0] = 0b00000000
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.Clean != false {
		t.Errorf("Clean has wrong value")
	}
	if packet.Spot != false {
		t.Errorf("Spot has wrong value")
	}
	if packet.Dock != false {
		t.Errorf("Dock has wrong value")
	}
	if packet.Minute != false {
		t.Errorf("Minute has wrong value")
	}
	if packet.Hour != false {
		t.Errorf("Hour has wrong value")
	}
	if packet.Day != false {
		t.Errorf("Day has wrong value")
	}
	if packet.Schedule != false {
		t.Errorf("Schedule has wrong value")
	}
	if packet.Clock != false {
		t.Errorf("Clock has wrong value")
	}
	data[0] = 0b11111111
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.Clean != true {
		t.Errorf("Clean has wrong value")
	}
	if packet.Spot != true {
		t.Errorf("Spot has wrong value")
	}
	if packet.Dock != true {
		t.Errorf("Dock has wrong value")
	}
	if packet.Minute != true {
		t.Errorf("Minute has wrong value")
	}
	if packet.Hour != true {
		t.Errorf("Hour has wrong value")
	}
	if packet.Day != true {
		t.Errorf("Day has wrong value")
	}
	if packet.Schedule != true {
		t.Errorf("Schedule has wrong value")
	}
	if packet.Clock != true {
		t.Errorf("Clock has wrong value")
	}
	data[0] = 0b01010101
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.Clean != true {
		t.Errorf("Clean has wrong value")
	}
	if packet.Spot != false {
		t.Errorf("Spot has wrong value")
	}
	if packet.Dock != true {
		t.Errorf("Dock has wrong value")
	}
	if packet.Minute != false {
		t.Errorf("Minute has wrong value")
	}
	if packet.Hour != true {
		t.Errorf("Hour has wrong value")
	}
	if packet.Day != false {
		t.Errorf("Day has wrong value")
	}
	if packet.Schedule != true {
		t.Errorf("Schedule has wrong value")
	}
	if packet.Clock != false {
		t.Errorf("Clock has wrong value")
	}
}
