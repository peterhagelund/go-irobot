package irobot

import "testing"

func TestExtract106(t *testing.T) {
	data := make([]byte, 12)
	packet := makePacket106().(*Packet106)
	err := packet.Extract(data, 0)
	if err != nil {
		t.Error(err)
	}
}
