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

// Packet32 is an encapsulation of a Packet with id 32.
type Packet32 struct {
	UnusedByte uint8
}

// ID returns the id.
func (packet *Packet32) ID() int {
	return 32
}

// Size returns the size.
func (packet *Packet32) Size() int {
	return 1
}

// Extract extracts the information in the specified data starting at the specified offset.
func (packet *Packet32) Extract(data []byte, offset int) error {
	if offset+packet.Size() > len(data) {
		return errors.New("packet exceeds data length")
	}
	packet.UnusedByte = data[offset]
	return nil
}

func newPacket32() Packet {
	return &Packet32{}
}
