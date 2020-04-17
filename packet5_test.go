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

func TestExtract5(t *testing.T) {
	packet := makePacket5().(*Packet5)
	data := []byte{
		2,    // Packet35
		4,    // Packet36
		1,    // Packet37
		42,   // Packet38
		0xfe, // Packet39
		0x0c, // -
		0xf8, // Packet40
		0x30, // -
		0x01, // Packet41
		0xf4, // -
		0xfe, // Packet42
		0x0c, // -
	}
	err := packet.Extract(data, 0)
	if err != nil {
		t.Fatal(err)
	}
	// Validate Packet35
	if packet.Packet35.Mode != ModeSafe {
		t.Fatalf("Mode has wrong value")
	}
	// Validate Packet36
	if packet.Packet36.Song != 4 {
		t.Fatalf("Song has wrong value")
	}
	// Validate Packet37
	if packet.Packet37.SongPlaying != true {
		t.Fatalf("SongPlaying has wrong value")
	}
	// Validate Packet38
	if packet.Packet38.StreamPacketCount != 42 {
		t.Fatalf("StreamPacketCount has wrong value")
	}
	// Validate Packet39
	if packet.Packet39.RequestedVelocity != -500 {
		t.Fatalf("RequestedVelocity has wrong value")
	}
	// Validate Packet40
	if packet.Packet40.RequestedRadius != -2000 {
		t.Fatalf("RequestedRadius has wrong value")
	}
	// Validate Packet41
	if packet.Packet41.RequestedRightVelocity != 500 {
		t.Fatalf("RequestedRightVelocity has wrong value")
	}
	// Validate Packet42
	if packet.Packet42.RequestedLeftVelocity != -500 {
		t.Fatalf("RequestedLeftVelocity has wrong value")
	}
}
