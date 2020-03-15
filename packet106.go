package irobot

import "errors"

// Packet106 is an encapsulation of a Packet with id 101.
type Packet106 struct {
}

// ID returns the id.
func (packet *Packet106) ID() int {
	return 106
}

// Size returns the size.
func (packet *Packet106) Size() int {
	return 12
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet106) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	return extractGroupPacket(packet, data, offset)
}

func makePacket106() Packet {
	return &Packet106{}
}
