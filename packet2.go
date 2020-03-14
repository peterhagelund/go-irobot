package irobot

import "errors"

// Packet2 is an encapsulation of a Packet with id 2.
type Packet2 struct {
	Packet17 *Packet17
	Packet18 *Packet18
	Packet19 *Packet19
	Packet20 *Packet20
}

// ID returns the id.
func (packet *Packet2) ID() int {
	return 2
}

// Size returns the size.
func (packet *Packet2) Size() int {
	return 6
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet2) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	return extractGroupPacket(packet, data, offset)
}

func makePacket2() Packet {
	return &Packet2{
		Packet17: &Packet17{},
		Packet18: &Packet18{},
		Packet19: &Packet19{},
		Packet20: &Packet20{},
	}
}
