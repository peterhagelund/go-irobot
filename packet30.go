package irobot

import "errors"

// Packet30 is an encapsulation of a Packet with id 30.
type Packet30 struct {
	CliffFrontRightSignal uint16
}

// ID returns the id.
func (packet *Packet30) ID() int {
	return 30
}

// Size returns the size.
func (packet *Packet30) Size() int {
	return 2
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet30) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	cliffFrontRightSignal := bytesToUint16(data[offset], data[offset+1])
	if cliffFrontRightSignal > 4095 {
		return errors.New("invalid cliff front right signal")
	}
	packet.CliffFrontRightSignal = cliffFrontRightSignal
	return nil
}

func makePacket30() Packet {
	return &Packet30{}
}
