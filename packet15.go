package irobot

import "errors"

// Packet15 is an encapsulation of a Packet with id 15.
type Packet15 struct {
	DirtDetect uint8
}

// ID returns the id.
func (packet *Packet15) ID() int {
	return 15
}

// Size returns the size.
func (packet *Packet15) Size() int {
	return 1
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet15) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	packet.DirtDetect = data[offset]
	return nil
}

func makePacket15() Packet {
	return &Packet15{}
}
