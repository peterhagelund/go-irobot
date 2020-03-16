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

// Packet4 is an encapsulation of a Packet with id 4.
type Packet4 struct {
	Packet27 *Packet27
	Packet28 *Packet28
	Packet29 *Packet29
	Packet30 *Packet30
	Packet31 *Packet31
	Packet32 *Packet32
	Packet33 *Packet33
	Packet34 *Packet34
}

// ID returns the id.
func (packet *Packet4) ID() int {
	return 4
}

// Size returns the size.
func (packet *Packet4) Size() int {
	return 14
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet4) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	return extractGroupPacket(packet, data, offset)
}

func makePacket4() Packet {
	return &Packet4{
		Packet27: &Packet27{},
		Packet28: &Packet28{},
		Packet29: &Packet29{},
		Packet30: &Packet30{},
		Packet31: &Packet31{},
		Packet32: &Packet32{},
		Packet33: &Packet33{},
		Packet34: &Packet34{},
	}
}
