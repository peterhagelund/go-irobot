package irobot

// Packet0 is an encapsulation of a Packet with id 0.
type Packet0 struct {
	Packet7  *Packet7
	Packet8  *Packet8
	Packet9  *Packet9
	Packet10 *Packet10
	Packet11 *Packet11
	Packet12 *Packet12
}

// ID returns the id.
func (packet *Packet0) ID() int {
	return 0
}

// Size returns the size.
func (packet *Packet0) Size() int {
	return 26
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet0) Extract(data []byte, offset int) error {
	return extractGroupPacket(packet, data, offset)
}

func makePacket0() Packet {
	return &Packet0{
		Packet7:  &Packet7{},
		Packet8:  &Packet8{},
		Packet9:  &Packet9{},
		Packet10: &Packet10{},
		Packet11: &Packet11{},
		Packet12: &Packet12{},
	}
}
