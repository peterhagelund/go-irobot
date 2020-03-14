package irobot

import "errors"

// Packet31 is an encapsulation of a Packet with id 31.
type Packet31 struct {
	CliffRightSignal uint16
}

// ID returns the id.
func (packet *Packet31) ID() int {
	return 31
}

// Size returns the size.
func (packet *Packet31) Size() int {
	return 2
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet31) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	cliffRightSignal := bytesToUint16(data[offset], data[offset+1])
	if cliffRightSignal > 4095 {
		return errors.New("invalid cliff left signal")
	}
	packet.CliffRightSignal = cliffRightSignal
	return nil
}

func makePacket31() Packet {
	return &Packet31{}
}
