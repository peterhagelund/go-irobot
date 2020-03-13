package irobot

import "errors"

// Packet19 is an encapsulation of a Packet with id 19.
type Packet19 struct {
	Distance int16
}

// ID returns the id.
func (packet *Packet19) ID() int {
	return 19
}

// Size returns the size.
func (packet *Packet19) Size() int {
	return 2
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet19) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	packet.Distance = bytesToInt16(data[offset], data[offset+1])
	return nil
}

func makePacket19() Packet {
	return &Packet19{}
}
