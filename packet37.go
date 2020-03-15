package irobot

import "errors"

// Packet37 is an encapsulation of a Packet with id 37.
type Packet37 struct {
	SongPlaying bool
}

// ID returns the id.
func (packet *Packet37) ID() int {
	return 37
}

// Size returns the size.
func (packet *Packet37) Size() int {
	return 1
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet37) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	packet.SongPlaying = data[offset] != 0x00
	return nil
}

func makePacket37() Packet {
	return &Packet37{}
}
