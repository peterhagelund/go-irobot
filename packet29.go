package irobot

import "errors"

// Packet29 is an encapsulation of a Packet with id 29.
type Packet29 struct {
	CliffFrontLeftSignal uint16
}

// ID returns the id.
func (packet *Packet29) ID() int {
	return 29
}

// Size returns the size.
func (packet *Packet29) Size() int {
	return 2
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet29) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	cliffFrontLeftSignal := bytesToUint16(data[offset], data[offset+1])
	if cliffFrontLeftSignal > 4095 {
		return errors.New("invalid cliff front left signal")
	}
	packet.CliffFrontLeftSignal = cliffFrontLeftSignal
	return nil
}

func makePacket29() Packet {
	return &Packet29{}
}
