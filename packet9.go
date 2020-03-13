package irobot

import "errors"

// Packet9 is an encapsulation of a Packet with id 9.
type Packet9 struct {
	CliffLeft bool
}

// ID returns the id.
func (packet *Packet9) ID() int {
	return 9
}

// Size returns the size.
func (packet *Packet9) Size() int {
	return 1
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet9) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	packet.CliffLeft = data[offset] != 0x00
	return nil
}

func makePacket9() Packet {
	return &Packet9{}
}
