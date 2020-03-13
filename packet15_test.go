package irobot

import "testing"

func TestExtract15(t *testing.T) {
	data := make([]byte, 1)
	packet := makePacket15().(*Packet15)
	data[0] = 0
	err := packet.Extract(data, 0)
	if err != nil {
		t.Error(err)
	}
	if packet.DirtDetect != 0 {
		t.Errorf("DirtDetect has wrong value")
	}
	data[0] = 42
	err = packet.Extract(data, 0)
	if err != nil {
		t.Error(err)
	}
	if packet.DirtDetect != 42 {
		t.Errorf("DirtDetect has wrong value")
	}
}
