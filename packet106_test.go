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

func TestExtract106(t *testing.T) {
	packet := makePacket106().(*Packet106)
	data := make([]byte, 12)
	data[0] = 0x03  // Packet46
	data[1] = 0x78  // -
	data[2] = 0x03  // Packet47
	data[3] = 0xe7  // -
	data[4] = 0x03  // Packet48
	data[5] = 0xf2  // -
	data[6] = 0x04  // Packet49
	data[7] = 0x57  // -
	data[8] = 0x04  // Packet50
	data[9] = 0xbc  // -
	data[10] = 0x05 // Packet51
	data[11] = 0x21 // -
	err := packet.Extract(data, 0)
	if err != nil {
		t.Error(err)
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
}
