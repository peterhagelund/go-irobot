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
	37:  makePacket37,
	38:  makePacket38,
	39:  makePacket39,
	40:  makePacket40,
	41:  makePacket41,
	42:  makePacket42,
	43:  makePacket43,
	44:  makePacket44,
	45:  makePacket45,
	46:  makePacket46,
	47:  makePacket47,
	48:  makePacket48,
	49:  makePacket49,
	50:  makePacket50,
	51:  makePacket51,
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
