package irobot

import "errors"

// Packet35 is an encapsulation of a Packet with id 35.
type Packet35 struct {
	Mode Mode
}

// ID returns the id.
func (packet *Packet35) ID() int {
	return 35
}

// Size returns the size.
func (packet *Packet35) Size() int {
	return 1
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet35) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	switch data[offset] {
	case 0:
		packet.Mode = ModeOff
	case 1:
		packet.Mode = ModePassive
	case 2:
		packet.Mode = ModeSafe
	case 3:
		packet.Mode = ModeFull
	default:
		packet.Mode = ModeUnknown
	}
	return nil
}

func makePacket35() Packet {
	return &Packet35{}
}
