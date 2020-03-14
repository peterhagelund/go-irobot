package irobot

import "errors"

// Packet4 is an encapsulation of a Packet with id 4.
type Packet4 struct {
	Packet27 *Packet27
	Packet28 *Packet28
	Packet29 *Packet29
	Packet30 *Packet30
	Packet31 *Packet31
}

// ID returns the id.
func (packet *Packet4) ID() int {
	return 4
}

// Size returns the size.
func (packet *Packet4) Size() int {
	return 14
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet4) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	return extractGroupPacket(packet, data, offset)
}

func makePacket4() Packet {
	return &Packet4{
		Packet27: &Packet27{},
		Packet28: &Packet28{},
		Packet29: &Packet29{},
		Packet30: &Packet30{},
		Packet31: &Packet31{},
	}
}
