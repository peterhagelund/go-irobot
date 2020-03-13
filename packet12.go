package irobot

import "errors"

// Packet12 is an encapsulation of a Packet with id 12.
type Packet12 struct {
	CliffRight bool
}

// ID returns the id.
func (packet *Packet12) ID() int {
	return 12
}

// Size returns the size.
func (packet *Packet12) Size() int {
	return 1
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet12) Extract(data []byte, offset int) error {
	if offset >= len(data) {
		return errors.New("offset exceeds data length")
	}
	packet.CliffRight = data[offset] != 0x00
	return nil
}

func makePacket12() Packet {
	return &Packet12{}
}
