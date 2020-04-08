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

import "testing"

func TestExtract34(t *testing.T) {
	packet := makePacket34().(*Packet34)
	data := make([]byte, 1)
	data[0] = 0b00000000
	if err := packet.Extract(data, 0); err != nil {
		t.Fatal(err)
	}
	if packet.InternalCharger != false {
		t.Fatal("InternalCharger has wrong value")
	}
	if packet.HomeBase != false {
		t.Fatal("HomeBase has wrong value")
	}
	data[0] = 0b00000011
	if err := packet.Extract(data, 0); err != nil {
		t.Fatal(err)
	}
	if packet.InternalCharger != true {
		t.Fatal("InternalCharger has wrong value")
	}
	if packet.HomeBase != true {
		t.Fatal("HomeBase has wrong value")
	}
	data[0] = 0b00000010
	if err := packet.Extract(data, 0); err != nil {
		t.Fatal(err)
	}
	if packet.InternalCharger != false {
		t.Fatal("InternalCharger has wrong value")
	}
	if packet.HomeBase != true {
		t.Fatal("HomeBase has wrong value")
	}
}
