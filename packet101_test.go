package irobot

import "testing"

func TestExtract101(t *testing.T) {
	packet := makePacket101().(*Packet101)
	data := make([]byte, 28)
	err := packet.Extract(data, 0)
	if err != nil {
		t.Error(err)
	}
}
