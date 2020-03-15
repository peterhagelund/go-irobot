package irobot

import "errors"

// Packet107 is an encapsulation of a Packet with id 101.
type Packet107 struct {
}

// ID returns the id.
func (packet *Packet107) ID() int {
	return 107
}

// Size returns the size.
func (packet *Packet107) Size() int {
	return 9
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet107) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	return extractGroupPacket(packet, data, offset)
}

func makePacket107() Packet {
	return &Packet107{}
}
