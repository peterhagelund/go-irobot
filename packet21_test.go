package irobot

import "testing"

func TestExtract21(t *testing.T) {
	packet := makePacket21().(*Packet21)
	data := make([]byte, 1)
	data[0] = 0
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.ChargingState != ChargingStateNotCharging {
		t.Errorf("ChargingState has wrong value")
	}
	data[0] = 5
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.ChargingState != ChargingStateChargingFaultCondition {
		t.Errorf("ChargingState has wrong value")
	}
	data[0] = 42
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.ChargingState != ChargingStateUnknown {
		t.Errorf("ChargingState has wrong value")
	}
}
