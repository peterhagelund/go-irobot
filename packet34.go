package irobot

import "errors"

// Packet34 is an encapsulation of a Packet with id 34.
type Packet34 struct {
	InternalCharger bool
	HomeBase        bool
}

// ID returns the id.
func (packet *Packet34) ID() int {
	return 34
}

// Size returns the size.
func (packet *Packet34) Size() int {
	return 1
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet34) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	bits := data[offset]
	packet.InternalCharger = bits&0b00000001 != 0b00000000
	packet.HomeBase = bits&0b00000010 != 0b00000000
	return nil
}

func makePacket34() Packet {
	return &Packet34{}
}
