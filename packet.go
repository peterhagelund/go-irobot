package irobot

// Packet defines the common, required behavior of a packet.
type Packet interface {
	// ID returns the packet ID.
	ID() uint8
	// Size returns the packet size.
	Size() uint8
	// Extract extracts the information in the specified data starting at the specified offset.
	Extract(data []byte, offset int) error
}

var packetFacory = map[uint8]func() Packet{
	0: makePacket0,
	7: makePacket7,
}

// Packet0 is an encapsulation of a Packet with id 0.
type Packet0 struct {
	packet7 *Packet7
}

// ID returns the id.
func (p *Packet0) ID() uint8 {
	return 0
}

// Extract extracts the information in the specified data starting at the specified offset.
func (p *Packet0) Extract(data []byte, offset int) error {
	return nil
}

// Size returns the size.
func (p *Packet0) Size() uint8 {
	return 26
}

// Packet7 is an encapsulation of a Packet with id 7.
type Packet7 struct {
	WheelDropLeft  bool
	WheelDropRight bool
	BumpLeft       bool
	BumpRight      bool
}

// ID returns the id.
func (p *Packet7) ID() uint8 {
	return 7
}

// Size returns the size.
func (p *Packet7) Size() uint8 {
	return 1
}

// Extract extracts the information in the specified data starting at the specified offset.
func (p *Packet7) Extract(data []byte, offset int) error {
	return nil
}

func makePacket0() Packet {
	return &Packet0{
		packet7: &Packet7{},
	}
}

func makePacket7() Packet {
	return &Packet7{}
}
