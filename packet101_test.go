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

func TestExtract101(t *testing.T) {
	packet := makePacket101().(*Packet101)
	data := make([]byte, 28)
	data[0] = 0x12       // Packet43
	data[1] = 0x34       // -
	data[2] = 0x23       // Packet44
	data[3] = 0x45       // -
	data[4] = 0b00010101 // Packet45
	data[5] = 0x03       // Packet46
	data[6] = 0x78       // -
	data[7] = 0x03       // Packet47
	data[8] = 0xe7       // -
	data[9] = 0x03       // Packet48
	data[10] = 0xf2      // -
	data[11] = 0x04      // Packet49
	data[12] = 0x57      // -
	data[13] = 0x04      // Packet50
	data[14] = 0xbc      // -
	data[15] = 0x05      // Packet51
	data[16] = 0x21      // -
	data[17] = 42        // Packet52
	data[18] = 13        // Packet53
	data[19] = 0xf8      // Packet54
	data[20] = 0x30      // -
	data[21] = 0x07      // Packet55
	data[22] = 0xd0      // -
	data[23] = 0x03      // Packet56
	data[24] = 0xe8      // -
	data[25] = 0xfc      // Packet57
	data[26] = 0x18      // -
	data[27] = 0x01      // Packet58
	err := packet.Extract(data, 0)
	if err != nil {
		t.Error(err)
	}
	// Validate Packet43
	if packet.Packet43.RightEncoderCounts != 4660 {
		t.Errorf("RightEncoderCounts has wrong value")
	}
	// Validate Packet44
	if packet.Packet44.LeftEncoderCounts != 9029 {
		t.Errorf("LeftEncoderCounts has wrong value")
	}
	// Validate Packet45
	if packet.Packet45.LightBumperLeft != true {
		t.Errorf("LightBumperLeft has wrong value")
	}
	if packet.Packet45.LightBumperFrontLeft != false {
		t.Errorf("LightBumperFrontLeft has wrong value")
	}
	if packet.Packet45.LightBumperCenterLeft != true {
		t.Errorf("LightBumperCenterLeft has wrong value")
	}
	if packet.Packet45.LightBumperCenterRight != false {
		t.Errorf("LightBumperCenterRight has wrong value")
	}
	if packet.Packet45.LightBumperFrontRight != true {
		t.Errorf("LightBumperFrontRight has wrong value")
	}
	if packet.Packet45.LightBumperRight != false {
		t.Errorf("LightBumperRight has wrong value")
	}
	// Validate Packet46
	if packet.Packet46.LightBumpLeftSignal != 888 {
		t.Errorf("LightBumpLeftSignal has wrong value")
	}
	// Validate Packet47
	if packet.Packet47.LightBumpFrontLeftSignal != 999 {
		t.Errorf("LightBumpFrontLeftSignal has wrong value")
	}
	// Validate Packet48
	if packet.Packet48.LightBumpCenterLeftSignal != 1010 {
		t.Errorf("LightBumpCenterLeftSignal has wrong value")
	}
	// Validate Packet49
	if packet.Packet49.LightBumpCenterRightSignal != 1111 {
		t.Errorf("LightBumpCenterRightSignal has wrong value")
	}
	// Validate Packet50
	if packet.Packet50.LightBumpFrontRightSignal != 1212 {
		t.Errorf("LightBumpFrontRightSignal has wrong value")
	}
	// Validate Packet51
	if packet.Packet51.LightBumpRightSignal != 1313 {
		t.Errorf("LightBumpRightSignal has wrong value")
	}
	// Validate Packet52
	if packet.Packet52.InfraredCharacterLeft != 42 {
		t.Errorf("InfraredCharacterLeft has wrong value")
	}
	// Validate Packet53
	if packet.Packet53.InfraredCharacterRight != 13 {
		t.Errorf("InfraredCharacterRight has wrong value")
	}
	// Validate Packet54
	if packet.Packet54.LeftMotorCurrent != -2000 {
		t.Errorf("LeftMotorCurrent has wrong value")
	}
	// Validate Packet55
	if packet.Packet55.RightMotorCurrent != 2000 {
		t.Errorf("RightMotorCurrent has wrong value")
	}
	// Validate Packet56
	if packet.Packet56.MainBrushMotorCurrent != 1000 {
		t.Errorf("MainBrushMotorCurrent has wrong value")
	}
	// Validate Packet57
	if packet.Packet57.SideBrushMotorCurrent != -1000 {
		t.Errorf("SideBrushMotorCurrent has wrong value")
	}
	// Validate Packet58
	if packet.Packet58.Stasis != true {
		t.Errorf("Stasis has wrong value")
	}
}
