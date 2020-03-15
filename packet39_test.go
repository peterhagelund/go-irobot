package irobot

import "testing"

func TestExtract39(t *testing.T) {
	packet := makePacket39().(*Packet39)
	data := make([]byte, 2)
	data[0] = 0x00
	data[1] = 0x00
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.RequestedVelocity != 0 {
		t.Errorf("RequestedVelocity has wrong value")
	}
	data[0] = 0xfe
	data[1] = 0x0c
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.RequestedVelocity != -500 {
		t.Errorf("RequestedVelocity has wrong value")
	}
	data[0] = 0x01
	data[1] = 0xf4
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.RequestedVelocity != 500 {
		t.Errorf("RequestedVelocity has wrong value")
	}
	data[0] = 0xfe
	data[1] = 0x0b
	if err := packet.Extract(data, 0); err == nil {
		t.Error("invalid requested velocity not rejected")
	}
	data[0] = 0x01
	data[1] = 0xf5
	if err := packet.Extract(data, 0); err == nil {
		t.Error("invalid requested velocity not rejected")
	}
}
