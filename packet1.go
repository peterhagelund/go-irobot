package irobot

import "errors"

// Packet1 is an encapsulation of a Packet with id 1.
type Packet1 struct {
	Packet7  *Packet7
	Packet8  *Packet8
	Packet9  *Packet9
	Packet10 *Packet10
	Packet11 *Packet11
	Packet12 *Packet12
	Packet13 *Packet13
	Packet14 *Packet14
	Packet15 *Packet15
	Packet16 *Packet16
}

// ID returns the id.
func (packet *Packet1) ID() int {
	return 1
}

// Size returns the size.
func (packet *Packet1) Size() int {
	return 10
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet1) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	return extractGroupPacket(packet, data, offset)
}

func makePacket1() Packet {
	return &Packet1{
		Packet7:  &Packet7{},
		Packet8:  &Packet8{},
		Packet9:  &Packet9{},
		Packet10: &Packet10{},
		Packet11: &Packet11{},
		Packet12: &Packet12{},
		Packet13: &Packet13{},
		Packet14: &Packet14{},
		Packet15: &Packet15{},
		Packet16: &Packet16{},
	}
}
