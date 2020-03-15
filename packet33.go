package irobot

import "errors"

// Packet33 is an encapsulation of a Packet with id 16.
type Packet33 struct {
	UnusedWord uint16
}

// ID returns the id.
func (packet *Packet33) ID() int {
	return 33
}

// Size returns the size.
func (packet *Packet33) Size() int {
	return 2
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet33) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	packet.UnusedWord = bytesToUint16(data[offset], data[offset+1])
	return nil
}

func makePacket33() Packet {
	return &Packet33{}
}
