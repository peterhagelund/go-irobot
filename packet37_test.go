package irobot

import "testing"

func TestExtract37(t *testing.T) {
	packet := makePacket37().(*Packet37)
	data := make([]byte, 1)
	data[0] = 0x00
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.SongPlaying != false {
		t.Errorf("SongPlaying has wrong value")
	}
	data[0] = 0x01
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.SongPlaying != true {
		t.Errorf("SongPlaying has wrong value")
	}
}
