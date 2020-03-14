package irobot

import "errors"

// Packet5 is an encapsulation of a Packet with id 5.
type Packet5 struct {
}

// ID returns the id.
func (packet *Packet5) ID() int {
	return 5
}

// Size returns the size.
func (packet *Packet5) Size() int {
	return 12
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet5) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	return extractGroupPacket(packet, data, offset)
}

func makePacket5() Packet {
	return &Packet5{}
}
