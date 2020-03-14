package irobot

import "errors"

// Packet28 is an encapsulation of a Packet with id 28.
type Packet28 struct {
	CliffLeftSignal uint16
}

// ID returns the id.
func (packet *Packet28) ID() int {
	return 28
}

// Size returns the size.
func (packet *Packet28) Size() int {
	return 2
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet28) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	cliffLeftSignal := bytesToUint16(data[offset], data[offset+1])
	if cliffLeftSignal > 4095 {
		return errors.New("invalid cliff left signal")
	}
	packet.CliffLeftSignal = cliffLeftSignal
	return nil
}

func makePacket28() Packet {
	return &Packet28{}
}
