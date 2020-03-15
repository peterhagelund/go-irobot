package irobot

import "testing"

func TestExtract38(t *testing.T) {
	packet := makePacket38().(*Packet38)
	data := make([]byte, 1)
	data[0] = 0
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.StreamPacketCount != 0 {
		t.Errorf("StreamPacketCount has wrong value")
	}
	data[0] = 108
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.StreamPacketCount != 108 {
		t.Errorf("StreamPacketCount has wrong value")
	}
	data[0] = 109
	if err := packet.Extract(data, 0); err == nil {
		t.Error("invalid stream packet count not rejected")
	}
}
