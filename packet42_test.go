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

func TestExtract42(t *testing.T) {
	packet := makePacket42().(*Packet42)
	data := make([]byte, 2)
	data[0] = 0x00
	data[1] = 0x00
	if err := packet.Extract(data, 0); err != nil {
		t.Fatal(err)
	}
	if packet.RequestedLeftVelocity != 0 {
		t.Fatal("RequestedLeftVelocity has wrong value")
	}
	data[0] = 0xfe
	data[1] = 0x0c
	if err := packet.Extract(data, 0); err != nil {
		t.Fatal(err)
	}
	if packet.RequestedLeftVelocity != -500 {
		t.Fatal("RequestedLeftVelocity has wrong value")
	}
	data[0] = 0x01
	data[1] = 0xf4
	if err := packet.Extract(data, 0); err != nil {
		t.Fatal(err)
	}
	if packet.RequestedLeftVelocity != 500 {
		t.Fatal("RequestedLeftVelocity has wrong value")
	}
	data[0] = 0xfe
	data[1] = 0x0b
	if err := packet.Extract(data, 0); err == nil {
		t.Fatal("invalid requested left velocity not rejected")
	}
	data[0] = 0x01
	data[1] = 0xf5
	if err := packet.Extract(data, 0); err == nil {
		t.Fatal("invalid requested left velocity not rejected")
	}
}
