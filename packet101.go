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

// Packet101 is an encapsulation of a Packet with id 101.
type Packet101 struct {
	Packet43 *Packet43
	Packet44 *Packet44
	Packet45 *Packet45
	Packet46 *Packet46
	Packet47 *Packet47
	Packet48 *Packet48
	Packet49 *Packet49
	Packet50 *Packet50
	Packet51 *Packet51
	Packet52 *Packet52
	Packet53 *Packet53
	Packet54 *Packet54
	Packet55 *Packet55
	Packet56 *Packet56
	Packet57 *Packet57
	Packet58 *Packet58
}

// ID returns the id.
func (packet *Packet101) ID() int {
	return 101
}

// Size returns the size.
func (packet *Packet101) Size() int {
	return 28
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet101) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	return extractGroupPacket(packet, data, offset)
}

func makePacket101() Packet {
	return &Packet101{
		Packet43: &Packet43{},
		Packet44: &Packet44{},
		Packet45: &Packet45{},
		Packet46: &Packet46{},
		Packet47: &Packet47{},
		Packet48: &Packet48{},
		Packet49: &Packet49{},
		Packet50: &Packet50{},
		Packet51: &Packet51{},
		Packet52: &Packet52{},
		Packet53: &Packet53{},
		Packet54: &Packet54{},
		Packet55: &Packet55{},
		Packet56: &Packet56{},
		Packet57: &Packet57{},
		Packet58: &Packet58{},
	}
}
