// Copyright (c) 2020-2024 Peter Hagelund
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

var packetFacory = map[string]func() Packet{
	"*irobot.Packet0":   newPacket0,
	"*irobot.Packet1":   newPacket1,
	"*irobot.Packet2":   newPacket2,
	"*irobot.Packet3":   newPacket3,
	"*irobot.Packet4":   newPacket4,
	"*irobot.Packet5":   newPacket5,
	"*irobot.Packet6":   newPacket6,
	"*irobot.Packet7":   newPacket7,
	"*irobot.Packet8":   newPacket8,
	"*irobot.Packet9":   newPacket9,
	"*irobot.Packet10":  newPacket10,
	"*irobot.Packet11":  newPacket11,
	"*irobot.Packet12":  newPacket12,
	"*irobot.Packet13":  newPacket13,
	"*irobot.Packet14":  newPacket14,
	"*irobot.Packet15":  newPacket15,
	"*irobot.Packet16":  newPacket16,
	"*irobot.Packet17":  newPacket17,
	"*irobot.Packet18":  newPacket18,
	"*irobot.Packet19":  newPacket19,
	"*irobot.Packet20":  newPacket20,
	"*irobot.Packet21":  newPacket21,
	"*irobot.Packet22":  newPacket22,
	"*irobot.Packet23":  newPacket23,
	"*irobot.Packet24":  newPacket24,
	"*irobot.Packet25":  newPacket25,
	"*irobot.Packet26":  newPacket26,
	"*irobot.Packet27":  newPacket27,
	"*irobot.Packet28":  newPacket28,
	"*irobot.Packet29":  newPacket29,
	"*irobot.Packet30":  newPacket30,
	"*irobot.Packet31":  newPacket31,
	"*irobot.Packet32":  newPacket32,
	"*irobot.Packet33":  newPacket33,
	"*irobot.Packet34":  newPacket34,
	"*irobot.Packet35":  newPacket35,
	"*irobot.Packet36":  newPacket36,
	"*irobot.Packet37":  newPacket37,
	"*irobot.Packet38":  newPacket38,
	"*irobot.Packet39":  newPacket39,
	"*irobot.Packet40":  newPacket40,
	"*irobot.Packet41":  newPacket41,
	"*irobot.Packet42":  newPacket42,
	"*irobot.Packet43":  newPacket43,
	"*irobot.Packet44":  newPacket44,
	"*irobot.Packet45":  newPacket45,
	"*irobot.Packet46":  newPacket46,
	"*irobot.Packet47":  newPacket47,
	"*irobot.Packet48":  newPacket48,
	"*irobot.Packet49":  newPacket49,
	"*irobot.Packet50":  newPacket50,
	"*irobot.Packet51":  newPacket51,
	"*irobot.Packet52":  newPacket52,
	"*irobot.Packet53":  newPacket53,
	"*irobot.Packet54":  newPacket54,
	"*irobot.Packet55":  newPacket55,
	"*irobot.Packet56":  newPacket56,
	"*irobot.Packet57":  newPacket57,
	"*irobot.Packet58":  newPacket58,
	"*irobot.Packet100": newPacket100,
	"*irobot.Packet101": newPacket101,
	"*irobot.Packet106": newPacket106,
	"*irobot.Packet107": newPacket107,
}

// NewPacket creates and returns a new packet of type T.
//
// Please note that "hack" used with reflection requires the packets to
// be in the irobot package and named Packet<id>.
func NewPacket[T Packet]() (T, error) {
	var p T
	typeName := reflect.TypeOf(p).String()
	newPacket, ok := packetFacory[typeName]
	if !ok {
		return p, fmt.Errorf("packet of type %s is unknown", typeName)
	}
	packet := newPacket()
	p = packet.(T)
	return p, nil
}

// NewPacketWithID creates and returns a new packet for the specified id.
func NewPacketWithID(id int) (Packet, error) {
	typeName := fmt.Sprintf("*irobot.Packet%d", id)
	newPacket, ok := packetFacory[typeName]
	if !ok {
		return nil, fmt.Errorf("packet with id %d is unknown", id)
	}
	return newPacket(), nil
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
