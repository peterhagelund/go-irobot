package irobot

import "testing"

func TestExtract101(t *testing.T) {
	data := make([]byte, 28)
	packet := makePacket101().(*Packet101)
	err := packet.Extract(data, 0)
	if err != nil {
		t.Error(err)
	}
}
