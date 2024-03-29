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

import "errors"

// Packet6 is an encapsulation of a Packet with id 6.
type Packet6 struct {
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
	Packet27 *Packet27
	Packet28 *Packet28
	Packet29 *Packet29
	Packet30 *Packet30
	Packet31 *Packet31
	Packet32 *Packet32
	Packet33 *Packet33
	Packet34 *Packet34
	Packet35 *Packet35
	Packet36 *Packet36
	Packet37 *Packet37
	Packet38 *Packet38
	Packet39 *Packet39
	Packet40 *Packet40
	Packet41 *Packet41
	Packet42 *Packet42
}

// ID returns the id.
func (packet *Packet6) ID() int {
	return 6
}

// Size returns the size.
func (packet *Packet6) Size() int {
	return 52
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet6) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	return extractGroupPacket(packet, data, offset)
}

func newPacket6() Packet {
	return &Packet6{
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
		Packet27: &Packet27{},
		Packet28: &Packet28{},
		Packet29: &Packet29{},
		Packet30: &Packet30{},
		Packet31: &Packet31{},
		Packet32: &Packet32{},
		Packet33: &Packet33{},
		Packet34: &Packet34{},
		Packet35: &Packet35{},
		Packet36: &Packet36{},
		Packet37: &Packet37{},
		Packet38: &Packet38{},
		Packet39: &Packet39{},
		Packet40: &Packet40{},
		Packet41: &Packet41{},
		Packet42: &Packet42{},
	}
}
