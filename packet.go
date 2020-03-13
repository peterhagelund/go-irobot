package irobot

import (
	"fmt"
	"reflect"
)

// Packet defines the common, required behavior of a packet.
type Packet interface {
	// ID returns the packet ID.
	ID() int
	// Size returns the packet size.
	Size() int
	// Extract extracts the information in the specified data starting at the specified offset.
	Extract(data []byte, offset int) error
}

var packetFacory = map[int]func() Packet{
	0:  makePacket0,
	7:  makePacket7,
	8:  makePacket8,
	9:  makePacket9,
	10: makePacket10,
	11: makePacket11,
}

// NewPacket creates and returns a new packet for the specified id.
func NewPacket(id int) (Packet, error) {
	f, ok := packetFacory[id]
	if !ok {
		return nil, fmt.Errorf("packet with id %d is unknown", id)
	}
	return f(), nil
}

func extractGroupPacket(packet Packet, data []byte, offset int) error {
	value := reflect.ValueOf(packet).Elem()
	for i := 0; i < value.NumField(); i++ {
		f := value.Field(i)
		p := f.Interface().(Packet)
		if err := p.Extract(data, offset); err != nil {
			return err
		}
		offset += p.Size()
	}
	return nil
}
