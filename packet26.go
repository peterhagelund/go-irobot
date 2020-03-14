package irobot

import "errors"

// Packet26 is an encapsulation of a Packet with id 26.
type Packet26 struct {
	BatteryCapacity uint16
}

// ID returns the id.
func (packet *Packet26) ID() int {
	return 26
}

// Size returns the size.
func (packet *Packet26) Size() int {
	return 2
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet26) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	packet.BatteryCapacity = bytesToUint16(data[offset], data[offset+1])
	return nil
}

func makePacket26() Packet {
	return &Packet26{}
}
