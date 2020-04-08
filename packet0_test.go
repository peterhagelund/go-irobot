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

func TestExtract0(t *testing.T) {
	packet := makePacket0().(*Packet0)
	data := make([]byte, 26)
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
	// Validate Packet21
	if packet.Packet21.ChargingState != ChargingStateTrickleCharging {
		t.Fatal("ChargingState has wrong value")
	}
	// Validate Packet22
	if packet.Packet22.Voltage != 16543 {
		t.Fatal("Voltage has wrong value")
	}
	// Validate Packet23
	if packet.Packet23.Current != -1234 {
		t.Fatal("Current has wrong value")
	}
	// Validate Packet24
	if packet.Packet24.Temperature != 19 {
		t.Fatal("Temperature has wrong value")
	}
	// Validate Packet25
	if packet.Packet25.BatteryCharge != 5000 {
		t.Fatal("BatteryCharge has wrong value")
	}
	// Validate Packet26
	if packet.Packet26.BatteryCapacity != 10000 {
		t.Fatal("BatteryCapacity has wrong value")
	}
}
