package irobot

import "errors"

// Packet10 is an encapsulation of a Packet with id 10.
type Packet10 struct {
	CliffFrontLeft bool
}

// ID returns the id.
func (packet *Packet10) ID() int {
	return 10
}

// Size returns the size.
func (packet *Packet10) Size() int {
	return 1
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet10) Extract(data []byte, offset int) error {
	if offset >= len(data) {
		return errors.New("offset exceeds data length")
	}
	packet.CliffFrontLeft = data[offset] != 0x00
	return nil
}

func makePacket10() Packet {
	return &Packet10{}
}
