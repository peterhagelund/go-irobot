package irobot

import (
	"fmt"
	"reflect"
	"testing"
)

var packetSizes = map[int]int{
	0:   26,
	1:   10,
	2:   6,
	3:   10,
	4:   14,
	5:   12,
	6:   52,
	100: 80,
	101: 28,
	106: 12,
	107: 9,
}

func TestNewPacket(t *testing.T) {
	for id := range packetFacory {
		packet, err := NewPacket(id)
		if err != nil {
			t.Error(err)
		}
		if packet.ID() != id {
			t.Errorf("expected packet with id %d, got %d", id, packet.ID())
		}
		size, ok := packetSizes[id]
		if !ok {
			size = 1
		}
		if packet.Size() != size {
			t.Errorf("expected packet with size %d, got %d", size, packet.Size())
		}
		expected := fmt.Sprintf("*irobot.Packet%d", id)
		actual := reflect.TypeOf(packet).String()
		if actual != expected {
			t.Errorf("expected packet with name '%s', got '%s'", expected, actual)
		}
	}
}
