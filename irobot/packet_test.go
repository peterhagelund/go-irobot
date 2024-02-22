// Copyright (c) 2020 Peter Hagelund
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

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
	41:  2,
	42:  2,
	43:  2,
	44:  2,
	46:  2,
	47:  2,
	48:  2,
	49:  2,
	50:  2,
	51:  2,
	54:  2,
	55:  2,
	56:  2,
	57:  2,
	100: 80,
	101: 28,
	106: 12,
	107: 9,
}

func TestNewPacket(t *testing.T) {
	for id := range packetFacory {
		packet, err := NewPacket(id)
		if err != nil {
			t.Fatal(err)
		}
		if packet.ID() != id {
			t.Fatalf("expected packet %T to have id %d, got %d", packet, id, packet.ID())
		}
		size, ok := packetSizes[id]
		if !ok {
			size = 1
		}
		if packet.Size() != size {
			t.Fatalf("expected packet %T to have size %d, got %d", packet, size, packet.Size())
		}
		expected := fmt.Sprintf("*irobot.Packet%d", id)
		actual := reflect.TypeOf(packet).String()
		if actual != expected {
			t.Fatalf("expected packet %T to have name '%s', got '%s'", packet, expected, actual)
		}
	}
}
