package irobot

import "errors"

// Packet27 is an encapsulation of a Packet with id 27.
type Packet27 struct {
	WallSignal uint16
}

// ID returns the id.
func (packet *Packet27) ID() int {
	return 27
}

// Size returns the size.
func (packet *Packet27) Size() int {
	return 2
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet27) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	wallSignal := bytesToUint16(data[offset], data[offset+1])
	if wallSignal > 1023 {
		return errors.New("invalid wall signal")
	}
	packet.WallSignal = wallSignal
	return nil
}

func makePacket27() Packet {
	return &Packet27{}
}
