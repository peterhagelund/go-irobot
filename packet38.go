package irobot

import "errors"

// Packet38 is an encapsulation of a Packet with id 38.
type Packet38 struct {
	StreamPacketCount uint8
}

// ID returns the id.
func (packet *Packet38) ID() int {
	return 38
}

// Size returns the size.
func (packet *Packet38) Size() int {
	return 1
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet38) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	streamPacketCount := data[offset]
	if streamPacketCount > 108 {
		return errors.New("invalid stream packet count")
	}
	packet.StreamPacketCount = streamPacketCount
	return nil
}

func makePacket38() Packet {
	return &Packet38{}
}
