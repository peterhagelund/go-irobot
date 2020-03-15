package irobot

import "errors"

// Packet32 is an encapsulation of a Packet with id 16.
type Packet32 struct {
	UnusedByte uint8
}

// ID returns the id.
func (packet *Packet32) ID() int {
	return 32
}

// Size returns the size.
func (packet *Packet32) Size() int {
	return 1
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet32) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	packet.UnusedByte = data[offset]
	return nil
}

func makePacket32() Packet {
	return &Packet32{}
}
