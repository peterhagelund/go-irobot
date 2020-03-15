package irobot

import "testing"

func TestExtract107(t *testing.T) {
	data := make([]byte, 9)
	packet := makePacket107().(*Packet107)
	err := packet.Extract(data, 0)
	if err != nil {
		t.Error(err)
	}
}
