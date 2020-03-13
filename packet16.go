package irobot

import "errors"

// Packet16 is an encapsulation of a Packet with id 16.
type Packet16 struct {
	UnusedByte uint8
}

// ID returns the id.
func (packet *Packet16) ID() int {
	return 16
}

// Size returns the size.
func (packet *Packet16) Size() int {
	return 1
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet16) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	packet.UnusedByte = data[offset]
	return nil
}

func makePacket16() Packet {
	return &Packet16{}
}
