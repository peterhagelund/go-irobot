package irobot

import "testing"

func TestExtract7(t *testing.T) {
	packet := makePacket7().(*Packet7)
	data := make([]byte, 1)
	data[0] = 0b00000000
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.BumpRight != false {
		t.Errorf("BumpRight has wrong value")
	}
	if packet.BumpLeft != false {
		t.Errorf("BumpLeft has wrong value")
	}
	if packet.WheelDropRight != false {
		t.Errorf("WheelDropRight has wrong value")
	}
	if packet.WheelDropLeft != false {
		t.Errorf("WheelDropLeft has wrong value")
	}
	data[0] = 0b00001111
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.BumpRight != true {
		t.Errorf("BumpRight has wrong value")
	}
	if packet.BumpLeft != true {
		t.Errorf("BumpLeft has wrong value")
	}
	if packet.WheelDropRight != true {
		t.Errorf("WheelDropRight has wrong value")
	}
	if packet.WheelDropLeft != true {
		t.Errorf("WheelDropLeft has wrong value")
	}
}
