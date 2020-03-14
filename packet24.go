package irobot

import "errors"

// Packet24 is an encapsulation of a Packet with id 23.
type Packet24 struct {
	Temperature uint8
}

// ID returns the id.
func (packet *Packet24) ID() int {
	return 24
}

// Size returns the size.
func (packet *Packet24) Size() int {
	return 1
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet24) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	packet.Temperature = data[offset]
	return nil
}

func makePacket24() Packet {
	return &Packet24{}
}
