package irobot

import "testing"

func TestExtract107(t *testing.T) {
	packet := makePacket107().(*Packet107)
	data := make([]byte, 9)
	err := packet.Extract(data, 0)
	if err != nil {
		t.Error(err)
	}
}
