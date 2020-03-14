package irobot

import "errors"

// Packet21 is an encapsulation of a Packet with id 21.
type Packet21 struct {
	ChargingState ChargingState
}

// ID returns the id.
func (packet *Packet21) ID() int {
	return 21
}

// Size returns the size.
func (packet *Packet21) Size() int {
	return 1
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet21) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	switch data[offset] {
	case 0:
		packet.ChargingState = ChargingStateNotCharging
	case 1:
		packet.ChargingState = ChargingStateReconditioningCharging
	case 2:
		packet.ChargingState = ChargingStateFullCharging
	case 3:
		packet.ChargingState = ChargingStateTrickleCharging
	case 4:
		packet.ChargingState = ChargingStateWaiting
	case 5:
		packet.ChargingState = ChargingStateChargingFaultCondition
	default:
		packet.ChargingState = ChargingStateUnknown
	}
	return nil
}

func makePacket21() Packet {
	return &Packet21{}
}
