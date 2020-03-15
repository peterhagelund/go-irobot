package irobot

import "errors"

// Packet39 is an encapsulation of a Packet with id 39.
type Packet39 struct {
	RequestedVelocity int16
}

// ID returns the id.
func (packet *Packet39) ID() int {
	return 39
}

// Size returns the size.
func (packet *Packet39) Size() int {
	return 2
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet39) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	requestedVelocity := bytesToInt16(data[offset], data[offset+1])
	if requestedVelocity < -500 || requestedVelocity > 500 {
		return errors.New("invalid requested velocity")
	}
	packet.RequestedVelocity = requestedVelocity
	return nil
}

func makePacket39() Packet {
	return &Packet39{}
}
