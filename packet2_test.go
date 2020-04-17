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

func TestExtract2(t *testing.T) {
	packet := makePacket2().(*Packet2)
	data := []byte{
		0x40,       // Packet17
		0b00000100, // Packet18
		0x30,       // Packet19
		0x39,       // -
		0xfc,       // Packet20
		0x19,       // -
	}
	err := packet.Extract(data, 0)
	if err != nil {
		t.Fatal(err)
	}
	// Validate Packet17
	if packet.Packet17.InfraredCharacterOmni != 0x40 {
		t.Fatal("InfraredCharacterOmni has wrong value")
	}
	// Validate Packet18
	if packet.Packet18.Clean != false {
		t.Fatal("Clean has wrong value")
	}
	if packet.Packet18.Spot != false {
		t.Fatal("Spot has wrong value")
	}
	if packet.Packet18.Dock != true {
		t.Fatal("Dock has wrong value")
	}
	if packet.Packet18.Minute != false {
		t.Fatal("Minute has wrong value")
	}
	if packet.Packet18.Hour != false {
		t.Fatal("Hour has wrong value")
	}
	if packet.Packet18.Day != false {
		t.Fatal("Day has wrong value")
	}
	if packet.Packet18.Schedule != false {
		t.Fatal("Schedule has wrong value")
	}
	if packet.Packet18.Clock != false {
		t.Fatal("Clock has wrong value")
	}
	// Validate Packet19
	if packet.Packet19.Distance != 12345 {
		t.Fatal("Distance has wrong value")
	}
	// Validate Packet20
	if packet.Packet20.Angle != -999 {
		t.Fatal("Angle has wrong value")
	}
}
