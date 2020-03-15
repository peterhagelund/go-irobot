package irobot

import "testing"

func TestExtract106(t *testing.T) {
	packet := makePacket106().(*Packet106)
	data := make([]byte, 12)
	err := packet.Extract(data, 0)
	if err != nil {
		t.Error(err)
	}
}
