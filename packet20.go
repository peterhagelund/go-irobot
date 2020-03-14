package irobot

import "errors"

// Packet20 is an encapsulation of a Packet with id 20.
type Packet20 struct {
	Angle int16
}

// ID returns the id.
func (packet *Packet20) ID() int {
	return 20
}

// Size returns the size.
func (packet *Packet20) Size() int {
	return 2
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet20) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	packet.Angle = bytesToInt16(data[offset], data[offset+1])
	return nil
}

func makePacket20() Packet {
	return &Packet20{}
}
