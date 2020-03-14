package irobot

import "errors"

// Packet25 is an encapsulation of a Packet with id 25.
type Packet25 struct {
	BatteryCharge uint16
}

// ID returns the id.
func (packet *Packet25) ID() int {
	return 25
}

// Size returns the size.
func (packet *Packet25) Size() int {
	return 2
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet25) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	packet.BatteryCharge = bytesToUint16(data[offset], data[offset+1])
	return nil
}

func makePacket25() Packet {
	return &Packet25{}
}
