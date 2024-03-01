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

func TestExtract101(t *testing.T) {
	packet := newPacket101().(*Packet101)
	data := []byte{
		0x12,       // Packet43
		0x34,       // -
		0x23,       // Packet44
		0x45,       // -
		0b00010101, // Packet45
		0x03,       // Packet46
		0x78,       // -
		0x03,       // Packet47
		0xe7,       // -
		0x03,       // Packet48
		0xf2,       // -
		0x04,       // Packet49
		0x57,       // -
		0x04,       // Packet50
		0xbc,       // -
		0x05,       // Packet51
		0x21,       // -
		42,         // Packet52
		13,         // Packet53
		0xf8,       // Packet54
		0x30,       // -
		0x07,       // Packet55
		0xd0,       // -
		0x03,       // Packet56
		0xe8,       // -
		0xfc,       // Packet57
		0x18,       // -
		0x01,       // Packet58
	}
	err := packet.Extract(data, 0)
	if err != nil {
		t.Fatal(err)
	}
	// Validate Packet43
	if packet.Packet43.RightEncoderCounts != 4660 {
		t.Fatal("RightEncoderCounts has wrong value")
	}
	// Validate Packet44
	if packet.Packet44.LeftEncoderCounts != 9029 {
		t.Fatal("LeftEncoderCounts has wrong value")
	}
	// Validate Packet45
	if packet.Packet45.LightBumperLeft != true {
		t.Fatal("LightBumperLeft has wrong value")
	}
	if packet.Packet45.LightBumperFrontLeft != false {
		t.Fatal("LightBumperFrontLeft has wrong value")
	}
	if packet.Packet45.LightBumperCenterLeft != true {
		t.Fatal("LightBumperCenterLeft has wrong value")
	}
	if packet.Packet45.LightBumperCenterRight != false {
		t.Fatal("LightBumperCenterRight has wrong value")
	}
	if packet.Packet45.LightBumperFrontRight != true {
		t.Fatal("LightBumperFrontRight has wrong value")
	}
	if packet.Packet45.LightBumperRight != false {
		t.Fatal("LightBumperRight has wrong value")
	}
	// Validate Packet46
	if packet.Packet46.LightBumpLeftSignal != 888 {
		t.Fatal("LightBumpLeftSignal has wrong value")
	}
	// Validate Packet47
	if packet.Packet47.LightBumpFrontLeftSignal != 999 {
		t.Fatal("LightBumpFrontLeftSignal has wrong value")
	}
	// Validate Packet48
	if packet.Packet48.LightBumpCenterLeftSignal != 1010 {
		t.Fatal("LightBumpCenterLeftSignal has wrong value")
	}
	// Validate Packet49
	if packet.Packet49.LightBumpCenterRightSignal != 1111 {
		t.Fatal("LightBumpCenterRightSignal has wrong value")
	}
	// Validate Packet50
	if packet.Packet50.LightBumpFrontRightSignal != 1212 {
		t.Fatal("LightBumpFrontRightSignal has wrong value")
	}
	// Validate Packet51
	if packet.Packet51.LightBumpRightSignal != 1313 {
		t.Fatal("LightBumpRightSignal has wrong value")
	}
	// Validate Packet52
	if packet.Packet52.InfraredCharacterLeft != 42 {
		t.Fatal("InfraredCharacterLeft has wrong value")
	}
	// Validate Packet53
	if packet.Packet53.InfraredCharacterRight != 13 {
		t.Fatal("InfraredCharacterRight has wrong value")
	}
	// Validate Packet54
	if packet.Packet54.LeftMotorCurrent != -2000 {
		t.Fatal("LeftMotorCurrent has wrong value")
	}
	// Validate Packet55
	if packet.Packet55.RightMotorCurrent != 2000 {
		t.Fatal("RightMotorCurrent has wrong value")
	}
	// Validate Packet56
	if packet.Packet56.MainBrushMotorCurrent != 1000 {
		t.Fatal("MainBrushMotorCurrent has wrong value")
	}
	// Validate Packet57
	if packet.Packet57.SideBrushMotorCurrent != -1000 {
		t.Fatal("SideBrushMotorCurrent has wrong value")
	}
	// Validate Packet58
	if packet.Packet58.Stasis != true {
		t.Fatal("Stasis has wrong value")
	}
}
