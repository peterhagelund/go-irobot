package irobot

import "errors"

// Packet11 is an encapsulation of a Packet with id 11.
type Packet11 struct {
	CliffFrontRight bool
}

// ID returns the id.
func (packet *Packet11) ID() int {
	return 11
}

// Size returns the size.
func (packet *Packet11) Size() int {
	return 1
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet11) Extract(data []byte, offset int) error {
	if offset >= len(data) {
		return errors.New("offset exceeds data length")
	}
	packet.CliffFrontRight = data[offset] != 0x00
	return nil
}

func makePacket11() Packet {
	return &Packet11{}
}
