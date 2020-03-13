package irobot

import "errors"

// Packet18 is an encapsulation of a Packet with id 18.
type Packet18 struct {
	Clean    bool
	Spot     bool
	Dock     bool
	Minute   bool
	Hour     bool
	Day      bool
	Schedule bool
	Clock    bool
}

// ID returns the id.
func (packet *Packet18) ID() int {
	return 18
}

// Size returns the size.
func (packet *Packet18) Size() int {
	return 1
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet18) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	bits := data[offset]
	packet.Clean = bits&0b00000001 != 0b00000000
	packet.Spot = bits&0b00000010 != 0b00000000
	packet.Dock = bits&0b00000100 != 0b00000000
	packet.Minute = bits&0b00001000 != 0b00000000
	packet.Hour = bits&0b00010000 != 0b00000000
	packet.Day = bits&0b00100000 != 0b00000000
	packet.Schedule = bits&0b01000000 != 0b00000000
	packet.Clock = bits&0b10000000 != 0b00000000
	return nil
}

func makePacket18() Packet {
	return &Packet18{}
}
