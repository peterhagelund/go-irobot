package irobot

import "errors"

// Packet101 is an encapsulation of a Packet with id 101.
type Packet101 struct {
}

// ID returns the id.
func (packet *Packet101) ID() int {
	return 101
}

// Size returns the size.
func (packet *Packet101) Size() int {
	return 28
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet101) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	return extractGroupPacket(packet, data, offset)
}

func makePacket101() Packet {
	return &Packet101{}
}
