package irobot

import "errors"

// Packet22 is an encapsulation of a Packet with id 22.
type Packet22 struct {
	Voltage uint16
}

// ID returns the id.
func (packet *Packet22) ID() int {
	return 22
}

// Size returns the size.
func (packet *Packet22) Size() int {
	return 2
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet22) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	packet.Voltage = bytesToUint16(data[offset], data[offset+1])
	return nil
}

func makePacket22() Packet {
	return &Packet22{}
}
