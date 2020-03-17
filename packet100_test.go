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

func TestExtract100(t *testing.T) {
	packet := makePacket100().(*Packet100)
	data := make([]byte, 80)
	data[0] = 0b00000010  // Packet7
	data[1] = 0x01        // Packet8
	data[2] = 0x01        // Packet9
	data[3] = 0x01        // Packet10
	data[4] = 0x00        // Packet11
	data[5] = 0x00        // Packet12
	data[6] = 0x01        // Packet13
	data[7] = 0b00011000  // Packet14
	data[8] = 42          // Packet15
	data[9] = 0x00        // Packet16
	data[10] = 0x40       // Packet17
	data[11] = 0b00000100 // Packet18
	data[12] = 0x30       // Packet19
	data[13] = 0x39       // -
	data[14] = 0xfc       // Packet20
	data[15] = 0x19       // -
	data[16] = 0x03       // Packet21
	data[17] = 0x40       // Packet22
	data[18] = 0x9f       // -
	data[19] = 0xfb       // Packet23
	data[20] = 0x2e       // -
	data[21] = 19         // Packet24
	data[22] = 0x13       // Packet25
	data[23] = 0x88       // -
	data[24] = 0x27       // Packet26
	data[25] = 0x10       // -
	data[26] = 0x03       // Packet27
	data[27] = 0xdb       // -
	data[28] = 0x0d       // Packet28
	data[29] = 0x80       // -
	data[30] = 0x09       // Packet29
	data[31] = 0x29       // -
	data[32] = 0x08       // Packet30
	data[33] = 0xae       // -
	data[34] = 0x04       // Packet31
	data[35] = 0x57       // -
	data[36] = 42         // Packet32
	data[37] = 0x12       // Packet33
	data[38] = 0x34       // -
	data[39] = 0b00000010 // Packet34
	data[40] = 2          // Packet35
	data[41] = 4          // Packet36
	data[42] = 1          // Packet37
	data[43] = 42         // Packet38
	data[44] = 0xfe       // Packet39
	data[45] = 0x0c       // -
	data[46] = 0xf8       // Packet40
	data[47] = 0x30       // -
	data[48] = 0x01       // Packet41
	data[49] = 0xf4       // -
	data[50] = 0xfe       // Packet42
	data[51] = 0x0c       // -
	data[52] = 0x12       // Packet43
	data[53] = 0x34       // -
	data[54] = 0x23       // Packet44
	data[55] = 0x45       // -
	data[56] = 0b00010101 // Packet45
	data[57] = 0x03       // Packet46
	data[58] = 0x78       // -
	data[59] = 0x03       // Packet47
	data[60] = 0xe7       // -
	data[61] = 0x03       // Packet48
	data[62] = 0xf2       // -
	data[63] = 0x04       // Packet49
	data[64] = 0x57       // -
	data[65] = 0x04       // Packet50
	data[66] = 0xbc       // -
	data[67] = 0x05       // Packet51
	data[68] = 0x21       // -
	data[69] = 42         // Packet52
	data[70] = 13         // Packet53
	data[71] = 0xf8       // Packet54
	data[72] = 0x30       // -
	data[73] = 0x07       // Packet55
	data[74] = 0xd0       // -
	data[75] = 0x03       // Packet56
	data[76] = 0xe8       // -
	data[77] = 0xfc       // Packet57
	data[78] = 0x18       // -
	data[79] = 0x01       // Packet58

	err := packet.Extract(data, 0)
	if err != nil {
		t.Error(err)
	}
	// Validate Packet7
	if packet.Packet7.BumpRight != false {
		t.Error("BumpRight has wrong value")
	}
	if packet.Packet7.BumpLeft != true {
		t.Error("BumpLeft has wrong value")
	}
	if packet.Packet7.WheelDropRight != false {
		t.Error("WheelDropRight has wrong value")
	}
	if packet.Packet7.WheelDropLeft != false {
		t.Error("WheelDropLeft has wrong value")
	}
	// Validate Packet8
	if packet.Packet8.Wall != true {
		t.Error("Wall has wrong value")
	}
	// Validate Packet9
	if packet.Packet9.CliffLeft != true {
		t.Error("CliffLeft has wrong value")
	}
	// Validate Packet10
	if packet.Packet10.CliffFrontLeft != true {
		t.Error("CliffFrontLeft has wrong value")
	}
	// Validate Packet11
	if packet.Packet11.CliffFrontRight != false {
		t.Error("CliffFrontRight has wrong value")
	}
	// Validate Packet12
	if packet.Packet12.CliffRight != false {
		t.Error("CliffRight has wrong value")
	}
	// Validate Packet13
	if packet.Packet13.VirtualWall != true {
		t.Error("VirtualWall has wrong value")
	}
	// Validate Packet14
	if packet.Packet14.SideBrush != false {
		t.Error("SideBrush has wrong value")
	}
	if packet.Packet14.MainBrush != false {
		t.Error("MainBrush has wrong value")
	}
	if packet.Packet14.RightWheel != true {
		t.Error("RightWheel has wrong value")
	}
	if packet.Packet14.LeftWheel != true {
		t.Error("LeftWheel has wrong value")
	}
	// Validate Packet15
	if packet.Packet15.DirtDetect != 42 {
		t.Error("DirtDetect has wrong value")
	}
	// Validate Packet16
	if packet.Packet16.UnusedByte != 0x00 {
		t.Error("UnusedByte has wrong value")
	}
	// Validate Packet17
	if packet.Packet17.InfraredCharacterOmni != 0x40 {
		t.Error("InfraredCharacterOmni has wrong value")
	}
	// Validate Packet18
	if packet.Packet18.Clean != false {
		t.Error("Clean has wrong value")
	}
	if packet.Packet18.Spot != false {
		t.Error("Spot has wrong value")
	}
	if packet.Packet18.Dock != true {
		t.Error("Dock has wrong value")
	}
	if packet.Packet18.Minute != false {
		t.Error("Minute has wrong value")
	}
	if packet.Packet18.Hour != false {
		t.Error("Hour has wrong value")
	}
	if packet.Packet18.Day != false {
		t.Error("Day has wrong value")
	}
	if packet.Packet18.Schedule != false {
		t.Error("Schedule has wrong value")
	}
	if packet.Packet18.Clock != false {
		t.Error("Clock has wrong value")
	}
	// Validate Packet19
	if packet.Packet19.Distance != 12345 {
		t.Error("Distance has wrong value")
	}
	// Validate Packet20
	if packet.Packet20.Angle != -999 {
		t.Error("Angle has wrong value")
	}
	// Validate Packet21
	if packet.Packet21.ChargingState != ChargingStateTrickleCharging {
		t.Error("ChargingState has wrong value")
	}
	// Validate Packet22
	if packet.Packet22.Voltage != 16543 {
		t.Error("Voltage has wrong value")
	}
	// Validate Packet23
	if packet.Packet23.Current != -1234 {
		t.Error("Current has wrong value")
	}
	// Validate Packet24
	if packet.Packet24.Temperature != 19 {
		t.Error("Temperature has wrong value")
	}
	// Validate Packet25
	if packet.Packet25.BatteryCharge != 5000 {
		t.Error("BatteryCharge has wrong value")
	}
	// Validate Packet26
	if packet.Packet26.BatteryCapacity != 10000 {
		t.Error("BatteryCapacity has wrong value")
	}
	// Validate Packet27
	if packet.Packet27.WallSignal != 987 {
		t.Error("WallSignal has wrong value")
	}
	// Validate Packet28
	if packet.Packet28.CliffLeftSignal != 3456 {
		t.Error("CliffLeftSignal has wrong value")
	}
	// Validate Packet29
	if packet.Packet29.CliffFrontLeftSignal != 2345 {
		t.Error("CliffFrontLeftSignal has wrong value")
	}
	// Validate Packet30
	if packet.Packet30.CliffFrontRightSignal != 2222 {
		t.Error("CliffFrontRightSignal has wrong value")
	}
	// Validate Packet31
	if packet.Packet31.CliffRightSignal != 1111 {
		t.Error("CliffFrontRightSignal has wrong value")
	}
	// Validate Packet32
	if packet.Packet32.UnusedByte != 42 {
		t.Errorf("UnusedByte has wrong value")
	}
	// Validate Packet33
	if packet.Packet33.UnusedWord != 4660 {
		t.Errorf("UnusedWord has wrong value")
	}
	// Validate Packet34
	if packet.Packet34.InternalCharger != false {
		t.Errorf("InternalCharger has wrong value")
	}
	if packet.Packet34.HomeBase != true {
		t.Errorf("HomeBase has wrong value")
	}
	// Validate Packet35
	if packet.Packet35.Mode != ModeSafe {
		t.Errorf("Mode has wrong value")
	}
	// Validate Packet36
	if packet.Packet36.Song != 4 {
		t.Errorf("Song has wrong value")
	}
	// Validate Packet37
	if packet.Packet37.SongPlaying != true {
		t.Errorf("SongPlaying has wrong value")
	}
	// Validate Packet38
	if packet.Packet38.StreamPacketCount != 42 {
		t.Errorf("StreamPacketCount has wrong value")
	}
	// Validate Packet39
	if packet.Packet39.RequestedVelocity != -500 {
		t.Errorf("RequestedVelocity has wrong value")
	}
	// Validate Packet40
	if packet.Packet40.RequestedRadius != -2000 {
		t.Errorf("RequestedRadius has wrong value")
	}
	// Validate Packet41
	if packet.Packet41.RequestedRightVelocity != 500 {
		t.Errorf("RequestedRightVelocity has wrong value")
	}
	// Validate Packet42
	if packet.Packet42.RequestedLeftVelocity != -500 {
		t.Errorf("RequestedLeftVelocity has wrong value")
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
