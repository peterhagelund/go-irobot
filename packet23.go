package irobot

import "errors"

// Packet23 is an encapsulation of a Packet with id 23.
type Packet23 struct {
	Current int16
}

// ID returns the id.
func (packet *Packet23) ID() int {
	return 23
}

// Size returns the size.
func (packet *Packet23) Size() int {
	return 2
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet23) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	packet.Current = bytesToInt16(data[offset], data[offset+1])
	return nil
}

func makePacket23() Packet {
	return &Packet23{}
}
