package irobot

import "errors"

// Packet40 is an encapsulation of a Packet with id 40.
type Packet40 struct {
	RequestedRadius int16
}

// ID returns the id.
func (packet *Packet40) ID() int {
	return 40
}

// Size returns the size.
func (packet *Packet40) Size() int {
	return 2
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet40) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	requestedRadius := bytesToInt16(data[offset], data[offset+1])
	if requestedRadius < -2000 || requestedRadius > 2000 {
		return errors.New("invalid requested radius")
	}
	packet.RequestedRadius = requestedRadius
	return nil
}

func makePacket40() Packet {
	return &Packet40{}
}
