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

import "testing"

func TestExtract54(t *testing.T) {
	packet := newPacket54().(*Packet54)
	data := []byte{
		0x00,
		0x00,
	}
	if err := packet.Extract(data, 0); err != nil {
		t.Fatal(err)
	}
	if packet.LeftMotorCurrent != 0 {
		t.Fatal("LeftMotorCurrent has wrong value")
	}
	data[0] = 0xf0
	data[1] = 0x60
	if err := packet.Extract(data, 0); err != nil {
		t.Fatal(err)
	}
	if packet.LeftMotorCurrent != -4000 {
		t.Fatal("LeftMotorCurrent has wrong value")
	}
	data[0] = 0x0f
	data[1] = 0xa0
	if err := packet.Extract(data, 0); err != nil {
		t.Fatal(err)
	}
	if packet.LeftMotorCurrent != 4000 {
		t.Fatal("LeftMotorCurrent has wrong value")
	}
}
