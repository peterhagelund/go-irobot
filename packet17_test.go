package irobot

import "testing"

func TestExtract17(t *testing.T) {
	packet := makePacket17().(*Packet17)
	data := make([]byte, 1)
	data[0] = 0
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.InfraredCharacterOmni != 0 {
		t.Errorf("InfraredCharacterOmni has wrong value")
	}
	data[0] = 42
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.InfraredCharacterOmni != 42 {
		t.Errorf("InfraredCharacterOmni has wrong value")
	}
}
