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

func TestExtract4(t *testing.T) {
	packet := makePacket4().(*Packet4)
	data := make([]byte, 14)
	data[0] = 0x03        // Packet27
	data[1] = 0xdb        // -
	data[2] = 0x0d        // Packet28
	data[3] = 0x80        // -
	data[4] = 0x09        // Packet29
	data[5] = 0x29        // -
	data[6] = 0x08        // Packet30
	data[7] = 0xae        // -
	data[8] = 0x04        // Packet31
	data[9] = 0x57        // -
	data[10] = 42         // Packet32
	data[11] = 0x12       // Packet33
	data[12] = 0x34       // -
	data[13] = 0b00000010 // Packet34
	err := packet.Extract(data, 0)
	if err != nil {
		t.Fatal(err)
	}
	// Validate Packet27
	if packet.Packet27.WallSignal != 987 {
		t.Fatal("WallSignal has wrong value")
	}
	// Validate Packet28
	if packet.Packet28.CliffLeftSignal != 3456 {
		t.Fatal("CliffLeftSignal has wrong value")
	}
	// Validate Packet29
	if packet.Packet29.CliffFrontLeftSignal != 2345 {
		t.Fatal("CliffFrontLeftSignal has wrong value")
	}
	// Validate Packet30
	if packet.Packet30.CliffFrontRightSignal != 2222 {
		t.Fatal("CliffFrontRightSignal has wrong value")
	}
	// Validate Packet31
	if packet.Packet31.CliffRightSignal != 1111 {
		t.Fatal("CliffFrontRightSignal has wrong value")
	}
	// Validate Packet32
	if packet.Packet32.UnusedByte != 42 {
		t.Fatalf("UnusedByte has wrong value")
	}
	// Validate Packet33
	if packet.Packet33.UnusedWord != 4660 {
		t.Fatalf("UnusedWord has wrong value")
	}
	// Validate Packet34
	if packet.Packet34.InternalCharger != false {
		t.Fatalf("InternalCharger has wrong value")
	}
	if packet.Packet34.HomeBase != true {
		t.Fatalf("HomeBase has wrong value")
	}
}
