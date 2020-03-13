package irobot

import "testing"

func TestBytesToUint16(t *testing.T) {
	if value := bytesToUint16(0x00, 0x00); value != 0 {
		t.Errorf("expected 0 but got %d", value)
	}
	if value := bytesToUint16(0x03, 0xe8); value != 1000 {
		t.Errorf("expected 1000 but got %d", value)
	}
	if value := bytesToUint16(0xff, 0xff); value != 65535 {
		t.Errorf("expected 65535 but got %d", value)
	}
}

func TestUint16ToBytes(t *testing.T) {
	if hi, lo := uint16ToBytes(0); hi != 0x00 || lo != 0x00 {
		t.Errorf("expected 0x00, 0x00 but got 0x%02x, 0x%02x", hi, lo)
	}
	if hi, lo := uint16ToBytes(1000); hi != 0x03 || lo != 0xe8 {
		t.Errorf("expected 0x03, 0xe8 but got 0x%02x, 0x%02x", hi, lo)
	}
	if hi, lo := uint16ToBytes(65535); hi != 0xff || lo != 0xff {
		t.Errorf("expected 0xff, 0xff but got 0x%02x, 0x%02x", hi, lo)
	}
}

func TestBytesToInt16(t *testing.T) {
	if value := bytesToInt16(0x00, 0x00); value != 0 {
		t.Errorf("expected 0 but got %d", value)
	}
	if value := bytesToInt16(0x01, 0xf4); value != 500 {
		t.Errorf("expected 500 but got %d", value)
	}
	if value := bytesToInt16(0x7f, 0xff); value != 32767 {
		t.Errorf("expected 32767 but got %d", value)
	}
	if value := bytesToInt16(0xff, 0x38); value != -200 {
		t.Errorf("expected -200 but got %d", value)
	}
	if value := bytesToInt16(0xfc, 0x18); value != -1000 {
		t.Errorf("expected -1000 but got %d", value)
	}
	if value := bytesToInt16(0xff, 0xff); value != -1 {
		t.Errorf("expected -1 but got %d", value)
	}
}

func TestIntToBytes(t *testing.T) {
	if hi, lo := int16ToBytes(0); hi != 0x00 || lo != 0x00 {
		t.Errorf("expected 0x00, 0x00 but got 0x%02x, 0x%02x", hi, lo)
	}
	if hi, lo := int16ToBytes(1000); hi != 0x03 || lo != 0xe8 {
		t.Errorf("expected 0x03, 0xe8 but got 0x%02x, 0x%02x", hi, lo)
	}
	if hi, lo := int16ToBytes(-200); hi != 0xff || lo != 0x38 {
		t.Errorf("expected 0xff, 0x38 but got 0x%02x, 0x%02x", hi, lo)
	}
	if hi, lo := int16ToBytes(-1000); hi != 0xfc || lo != 0x18 {
		t.Errorf("expected 0xfc, 0x18 but got 0x%02x, 0x%02x", hi, lo)
	}
	if hi, lo := int16ToBytes(-1); hi != 0xff || lo != 0xff {
		t.Errorf("expected 0xff, 0xff but got 0x%02x, 0x%02x", hi, lo)
	}
}
