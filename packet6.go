package irobot

import "errors"

// Packet6 is an encapsulation of a Packet with id 6.
type Packet6 struct {
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
	Packet17 *Packet17
	Packet18 *Packet18
	Packet19 *Packet19
	Packet20 *Packet20
	Packet21 *Packet21
	Packet22 *Packet22
	Packet23 *Packet23
	Packet24 *Packet24
	Packet25 *Packet25
	Packet26 *Packet26
	Packet27 *Packet27
	Packet28 *Packet28
	Packet29 *Packet29
	Packet30 *Packet30
	Packet31 *Packet31
}

// ID returns the id.
func (packet *Packet6) ID() int {
	return 6
}

// Size returns the size.
func (packet *Packet6) Size() int {
	return 52
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet6) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	return extractGroupPacket(packet, data, offset)
}

func makePacket6() Packet {
	return &Packet6{
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
		Packet17: &Packet17{},
		Packet18: &Packet18{},
		Packet19: &Packet19{},
		Packet20: &Packet20{},
		Packet21: &Packet21{},
		Packet22: &Packet22{},
		Packet23: &Packet23{},
		Packet24: &Packet24{},
		Packet25: &Packet25{},
		Packet26: &Packet26{},
		Packet27: &Packet27{},
		Packet28: &Packet28{},
		Packet29: &Packet29{},
		Packet30: &Packet30{},
		Packet31: &Packet31{},
	}
}
