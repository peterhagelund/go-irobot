package irobot

import "errors"

// Packet13 is an encapsulation of a Packet with id 13.
type Packet13 struct {
	VirtualWall bool
}

// ID returns the id.
func (packet *Packet13) ID() int {
	return 13
}

// Size returns the size.
func (packet *Packet13) Size() int {
	return 1
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet13) Extract(data []byte, offset int) error {
	if offset >= len(data) {
		return errors.New("offset exceeds data length")
	}
	packet.VirtualWall = data[offset] != 0x00
	return nil
}

func makePacket13() Packet {
	return &Packet13{}
}
