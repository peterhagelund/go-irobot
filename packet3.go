package irobot

import "errors"

// Packet3 is an encapsulation of a Packet with id 3.
type Packet3 struct {
	Packet21 *Packet21
	Packet22 *Packet22
	Packet23 *Packet23
	Packet24 *Packet24
	Packet25 *Packet25
	Packet26 *Packet26
}

// ID returns the id.
func (packet *Packet3) ID() int {
	return 3
}

// Size returns the size.
func (packet *Packet3) Size() int {
	return 10
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet3) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	return extractGroupPacket(packet, data, offset)
}

func makePacket3() Packet {
	return &Packet3{
		Packet21: &Packet21{},
		Packet22: &Packet22{},
		Packet23: &Packet23{},
		Packet24: &Packet24{},
		Packet25: &Packet25{},
		Packet26: &Packet26{},
	}
}
