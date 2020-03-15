package irobot

import "errors"

// Packet36 is an encapsulation of a Packet with id 36.
type Packet36 struct {
	Song uint8
}

// ID returns the id.
func (packet *Packet36) ID() int {
	return 36
}

// Size returns the size.
func (packet *Packet36) Size() int {
	return 1
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet36) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	song := data[offset]
	if song > 15 {
		return errors.New("invalid song number")
	}
	packet.Song = song
	return nil
}

func makePacket36() Packet {
	return &Packet36{}
}
