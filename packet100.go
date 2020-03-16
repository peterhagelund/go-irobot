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

import "errors"

// Packet100 is an encapsulation of a Packet with id 100.
type Packet100 struct {
	Packet7  *Packet7
	Packet8  *Packet8
	Packet9  *Packet9
	Packet10 *Packet10
	Packet11 *Packet11
	Packet12 *Packet12
	Packet13 *Packet13
	Packet14 *Packet14
	Packet15 *Packet15
	Packet16 *Packet16
	Packet17 *Packet17
	Packet18 *Packet18
	Packet19 *Packet19
	Packet20 *Packet20
	Packet21 *Packet21
	Packet22 *Packet22
	Packet23 *Packet23
	Packet24 *Packet24
	Packet25 *Packet25
	Packet26 *Packet26
}

// ID returns the id.
func (packet *Packet100) ID() int {
	return 100
}

// Size returns the size.
func (packet *Packet100) Size() int {
	return 80
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet100) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	return extractGroupPacket(packet, data, offset)
}

func makePacket100() Packet {
	return &Packet100{
		Packet7:  &Packet7{},
		Packet8:  &Packet8{},
		Packet9:  &Packet9{},
		Packet10: &Packet10{},
		Packet11: &Packet11{},
		Packet12: &Packet12{},
		Packet13: &Packet13{},
		Packet14: &Packet14{},
		Packet15: &Packet15{},
		Packet16: &Packet16{},
		Packet17: &Packet17{},
		Packet18: &Packet18{},
		Packet19: &Packet19{},
		Packet20: &Packet20{},
		Packet21: &Packet21{},
		Packet22: &Packet22{},
		Packet23: &Packet23{},
		Packet24: &Packet24{},
		Packet25: &Packet25{},
		Packet26: &Packet26{},
	}
}
