package irobot

import "errors"

// Packet8 is an encapsulation of a Packet with id 8.
type Packet8 struct {
	Wall bool
}

// ID returns the id.
func (packet *Packet8) ID() int {
	return 8
}

// Size returns the size.
func (packet *Packet8) Size() int {
	return 1
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet8) Extract(data []byte, offset int) error {
	if offset >= len(data) {
		return errors.New("offset exceeds data length")
	}
	packet.Wall = data[offset] != 0x00
	return nil
}

func makePacket8() Packet {
	return &Packet8{}
}
