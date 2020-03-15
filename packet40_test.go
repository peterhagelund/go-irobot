package irobot

import "testing"

func TestExtract40(t *testing.T) {
	packet := makePacket40().(*Packet40)
	data := make([]byte, 2)
	data[0] = 0x00
	data[1] = 0x00
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.RequestedRadius != 0 {
		t.Errorf("RequestedRadius has wrong value")
	}
	data[0] = 0xf8
	data[1] = 0x30
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.RequestedRadius != -2000 {
		t.Errorf("RequestedRadius has wrong value")
	}
	data[0] = 0x07
	data[1] = 0xd0
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.RequestedRadius != 2000 {
		t.Errorf("RequestedRadius has wrong value")
	}
	data[0] = 0xf8
	data[1] = 0x2f
	if err := packet.Extract(data, 0); err == nil {
		t.Error("invalid requested radius not rejected")
	}
	data[0] = 0x07
	data[1] = 0xd1
	if err := packet.Extract(data, 0); err == nil {
		t.Error("invalid requested radius not rejected")
	}
}
