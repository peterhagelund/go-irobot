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
	0:   makePacket0,
	1:   makePacket1,
	2:   makePacket2,
	3:   makePacket3,
	4:   makePacket4,
	5:   makePacket5,
	6:   makePacket6,
	7:   makePacket7,
	8:   makePacket8,
	9:   makePacket9,
	10:  makePacket10,
	11:  makePacket11,
	12:  makePacket12,
	13:  makePacket13,
	14:  makePacket14,
	15:  makePacket15,
	16:  makePacket16,
	17:  makePacket17,
	18:  makePacket18,
	19:  makePacket19,
	20:  makePacket20,
	21:  makePacket21,
	22:  makePacket22,
	23:  makePacket23,
	24:  makePacket24,
	25:  makePacket25,
	26:  makePacket26,
	27:  makePacket27,
	28:  makePacket28,
	29:  makePacket29,
	30:  makePacket30,
	31:  makePacket31,
	32:  makePacket32,
	33:  makePacket33,
	34:  makePacket34,
	35:  makePacket35,
	36:  makePacket36,
	100: makePacket100,
	101: makePacket101,
	106: makePacket106,
	107: makePacket107,
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
