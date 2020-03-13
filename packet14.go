package irobot

import "errors"

// Packet14 is an encapsulation of a Packet with id 14.
type Packet14 struct {
	SideBrush  bool
	MainBrush  bool
	RightWheel bool
	LeftWheel  bool
}

// ID returns the id.
func (packet *Packet14) ID() int {
	return 14
}

// Size returns the size.
func (packet *Packet14) Size() int {
	return 1
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet14) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	bits := data[offset]
	packet.SideBrush = bits&0b00000001 != 0b00000000
	packet.MainBrush = bits&0b00000100 != 0b00000000
	packet.RightWheel = bits&0b00001000 != 0b00000000
	packet.LeftWheel = bits&0b00010000 != 0b00000000
	return nil
}

func makePacket14() Packet {
	return &Packet14{}
}
