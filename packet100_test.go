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
}
