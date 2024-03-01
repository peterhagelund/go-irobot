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

// Packet7 is an encapsulation of a Packet with id 7.
type Packet7 struct {
	BumpRight      bool
	BumpLeft       bool
	WheelDropRight bool
	WheelDropLeft  bool
}

// ID returns the id.
func (packet *Packet7) ID() int {
	return 7
}

// Size returns the size.
func (packet *Packet7) Size() int {
	return 1
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet7) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	bits := data[offset]
	packet.BumpRight = bits&0b00000001 != 0b00000000
	packet.BumpLeft = bits&0b00000010 != 0b00000000
	packet.WheelDropRight = bits&0b00000100 != 0b00000000
	packet.WheelDropLeft = bits&0b00001000 != 0b00000000
	return nil
}

func newPacket7() Packet {
	return &Packet7{}
}
