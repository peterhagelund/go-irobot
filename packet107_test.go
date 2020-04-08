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

func TestExtract107(t *testing.T) {
	packet := makePacket107().(*Packet107)
	data := make([]byte, 9)
	data[0] = 0xf0 // Packet54
	data[1] = 0x60 // -
	data[2] = 0x0f // Packet55
	data[3] = 0xa0 // -
	data[4] = 0xf4 // Packet56
	data[5] = 0x48 // -
	data[6] = 0x0b // Packet57
	data[7] = 0xb8 // -
	data[8] = 0x01 // Packet58
	err := packet.Extract(data, 0)
	if err != nil {
		t.Fatal(err)
	}
	// Validate Packet54
	if packet.Packet54.LeftMotorCurrent != -4000 {
		t.Fatal("LeftMotorCurrent has wrong value")
	}
	// Validate Packet55
	if packet.Packet55.RightMotorCurrent != 4000 {
		t.Fatal("RightMotorCurrent has wrong value")
	}
	// Validate Packet56
	if packet.Packet56.MainBrushMotorCurrent != -3000 {
		t.Fatal("MainBrushMotorCurrent has wrong value")
	}
	// Validate Packet57
	if packet.Packet57.SideBrushMotorCurrent != 3000 {
		t.Fatal("SideBrushMotorCurrent has wrong value")
	}
	// Validate Packet58
	if packet.Packet58.Stasis != true {
		t.Fatal("Stasis has wrong value")
	}
}
