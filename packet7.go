package irobot

import "errors"

// Packet7 is an encapsulation of a Packet with id 7.
type Packet7 struct {
	BumpRight      bool
	BumpLeft       bool
	WheelDropRight bool
	WheelDropLeft  bool
}

// ID returns the id.
func (packet *Packet7) ID() int {
	return 7
}

// Size returns the size.
func (packet *Packet7) Size() int {
	return 1
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet7) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	bits := data[offset]
	packet.BumpRight = bits&0b00000001 != 0b00000000
	packet.BumpLeft = bits&0b00000010 != 0b00000000
	packet.WheelDropRight = bits&0b00000100 != 0b00000000
	packet.WheelDropLeft = bits&0b00001000 != 0b00000000
	return nil
}

func makePacket7() Packet {
	return &Packet7{}
}
