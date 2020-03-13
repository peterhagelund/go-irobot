package irobot

import "testing"

func TestExtract0(t *testing.T) {
	data := make([]byte, 26)
	packet := makePacket0().(*Packet0)
	err := packet.Extract(data, 0)
	if err != nil {
		t.Error(err)
	}
}
