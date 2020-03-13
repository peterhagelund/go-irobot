package irobot

import "testing"

func TestExtract14(t *testing.T) {
	data := make([]byte, 1)
	packet := makePacket14().(*Packet14)
	data[0] = 0b00000000
	err := packet.Extract(data, 0)
	if err != nil {
		t.Error(err)
	}
	if packet.SideBrush != false {
		t.Errorf("SideBrush has wrong value")
	}
	if packet.MainBrush != false {
		t.Errorf("MainBrush has wrong value")
	}
	if packet.RightWheel != false {
		t.Errorf("RightWheel has wrong value")
	}
	if packet.LeftWheel != false {
		t.Errorf("LeftWheel has wrong value")
	}
	data[0] = 0b00011101
	err = packet.Extract(data, 0)
	if err != nil {
		t.Error(err)
	}
	if packet.SideBrush != true {
		t.Errorf("SideBrush has wrong value")
	}
	if packet.MainBrush != true {
		t.Errorf("MainBrush has wrong value")
	}
	if packet.RightWheel != true {
		t.Errorf("RightWheel has wrong value")
	}
	if packet.LeftWheel != true {
		t.Errorf("LeftWheel has wrong value")
	}
}
