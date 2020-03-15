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
	19:  2,
	20:  2,
	22:  2,
	23:  2,
	25:  2,
	26:  2,
	27:  2,
	28:  2,
	29:  2,
	30:  2,
	31:  2,
	33:  2,
	39:  2,
	40:  2,
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
			t.Errorf("expected packet %T to have id %d, got %d", packet, id, packet.ID())
		}
		size, ok := packetSizes[id]
		if !ok {
			size = 1
		}
		if packet.Size() != size {
			t.Errorf("expected packet %T to have size %d, got %d", packet, size, packet.Size())
		}
		expected := fmt.Sprintf("*irobot.Packet%d", id)
		actual := reflect.TypeOf(packet).String()
		if actual != expected {
			t.Errorf("expected packet %T to have name '%s', got '%s'", packet, expected, actual)
		}
	}
}
