package irobot

import "errors"

// Packet17 is an encapsulation of a Packet with id 17.
type Packet17 struct {
	InfraredCharacterOmni uint8
}

// ID returns the id.
func (packet *Packet17) ID() int {
	return 17
}

// Size returns the size.
func (packet *Packet17) Size() int {
	return 1
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet17) Extract(data []byte, offset int) error {
	if offset >= len(data) {
		return errors.New("offset exceeds data length")
	}
	packet.InfraredCharacterOmni = data[offset]
	return nil
}

func makePacket17() Packet {
	return &Packet17{}
}
