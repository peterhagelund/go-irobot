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

func TestExtract1(t *testing.T) {
	packet := makePacket1().(*Packet1)
	data := make([]byte, 10)
	data[0] = 0b00000010 // Packet7
	data[1] = 0x01       // Packet8
	data[2] = 0x01       // Packet9
	data[3] = 0x01       // Packet10
	data[4] = 0x00       // Packet11
	data[5] = 0x00       // Packet12
	data[6] = 0x01       // Packet13
	data[7] = 0b00011000 // Packet14
	data[8] = 42         // Packet15
	data[9] = 0x00       // Packet16
	err := packet.Extract(data, 0)
	if err != nil {
		t.Fatal(err)
	}
	// Validate Packet7
	if packet.Packet7.BumpRight != false {
		t.Fatal("BumpRight has wrong value")
	}
	if packet.Packet7.BumpLeft != true {
		t.Fatal("BumpLeft has wrong value")
	}
	if packet.Packet7.WheelDropRight != false {
		t.Fatal("WheelDropRight has wrong value")
	}
	if packet.Packet7.WheelDropLeft != false {
		t.Fatal("WheelDropLeft has wrong value")
	}
	// Validate Packet8
	if packet.Packet8.Wall != true {
		t.Fatal("Wall has wrong value")
	}
	// Validate Packet9
	if packet.Packet9.CliffLeft != true {
		t.Fatal("CliffLeft has wrong value")
	}
	// Validate Packet10
	if packet.Packet10.CliffFrontLeft != true {
		t.Fatal("CliffFrontLeft has wrong value")
	}
	// Validate Packet11
	if packet.Packet11.CliffFrontRight != false {
		t.Fatal("CliffFrontRight has wrong value")
	}
	// Validate Packet12
	if packet.Packet12.CliffRight != false {
		t.Fatal("CliffRight has wrong value")
	}
	// Validate Packet13
	if packet.Packet13.VirtualWall != true {
		t.Fatal("VirtualWall has wrong value")
	}
	// Validate Packet14
	if packet.Packet14.SideBrush != false {
		t.Fatal("SideBrush has wrong value")
	}
	if packet.Packet14.MainBrush != false {
		t.Fatal("MainBrush has wrong value")
	}
	if packet.Packet14.RightWheel != true {
		t.Fatal("RightWheel has wrong value")
	}
	if packet.Packet14.LeftWheel != true {
		t.Fatal("LeftWheel has wrong value")
	}
	// Validate Packet15
	if packet.Packet15.DirtDetect != 42 {
		t.Fatal("DirtDetect has wrong value")
	}
	// Validate Packet16
	if packet.Packet16.UnusedByte != 0x00 {
		t.Fatal("UnusedByte has wrong value")
	}
}
