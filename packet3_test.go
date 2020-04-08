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

func TestExtract3(t *testing.T) {
	packet := makePacket3().(*Packet3)
	data := make([]byte, 10)
	data[0] = 0x03 // Packet21
	data[1] = 0x40 // Packet22
	data[2] = 0x9f // -
	data[3] = 0xfb // Packet23
	data[4] = 0x2e // -
	data[5] = 19   // Packet24
	data[6] = 0x13 // Packet25
	data[7] = 0x88 // -
	data[8] = 0x27 // Packet26
	data[9] = 0x10 // -
	err := packet.Extract(data, 0)
	if err != nil {
		t.Fatal(err)
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
