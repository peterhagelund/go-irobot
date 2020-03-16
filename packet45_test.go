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

func TestExtract45(t *testing.T) {
	packet := makePacket45().(*Packet45)
	data := make([]byte, 1)
	data[0] = 0b00000000
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.LightBumperLeft != false {
		t.Errorf("LightBumperLeft has wrong value")
	}
	if packet.LightBumperFrontLeft != false {
		t.Errorf("LightBumperFrontLeft has wrong value")
	}
	if packet.LightBumperCenterLeft != false {
		t.Errorf("LightBumperCenterLeft has wrong value")
	}
	if packet.LightBumperCenterRight != false {
		t.Errorf("LightBumperCenterRight has wrong value")
	}
	if packet.LightBumperFrontRight != false {
		t.Errorf("LightBumperFrontRight has wrong value")
	}
	if packet.LightBumperRight != false {
		t.Errorf("LightBumperRight has wrong value")
	}
	data[0] = 0b00010101
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.LightBumperLeft != true {
		t.Errorf("LightBumperLeft has wrong value")
	}
	if packet.LightBumperFrontLeft != false {
		t.Errorf("LightBumperFrontLeft has wrong value")
	}
	if packet.LightBumperCenterLeft != true {
		t.Errorf("LightBumperCenterLeft has wrong value")
	}
	if packet.LightBumperCenterRight != false {
		t.Errorf("LightBumperCenterRight has wrong value")
	}
	if packet.LightBumperFrontRight != true {
		t.Errorf("LightBumperFrontRight has wrong value")
	}
	if packet.LightBumperRight != false {
		t.Errorf("LightBumperRight has wrong value")
	}
	data[0] = 0b00101010
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.LightBumperLeft != false {
		t.Errorf("LightBumperLeft has wrong value")
	}
	if packet.LightBumperFrontLeft != true {
		t.Errorf("LightBumperFrontLeft has wrong value")
	}
	if packet.LightBumperCenterLeft != false {
		t.Errorf("LightBumperCenterLeft has wrong value")
	}
	if packet.LightBumperCenterRight != true {
		t.Errorf("LightBumperCenterRight has wrong value")
	}
	if packet.LightBumperFrontRight != false {
		t.Errorf("LightBumperFrontRight has wrong value")
	}
	if packet.LightBumperRight != true {
		t.Errorf("LightBumperRight has wrong value")
	}
	data[0] = 0b00111111
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.LightBumperLeft != true {
		t.Errorf("LightBumperLeft has wrong value")
	}
	if packet.LightBumperFrontLeft != true {
		t.Errorf("LightBumperFrontLeft has wrong value")
	}
	if packet.LightBumperCenterLeft != true {
		t.Errorf("LightBumperCenterLeft has wrong value")
	}
	if packet.LightBumperCenterRight != true {
		t.Errorf("LightBumperCenterRight has wrong value")
	}
	if packet.LightBumperFrontRight != true {
		t.Errorf("LightBumperFrontRight has wrong value")
	}
	if packet.LightBumperRight != true {
		t.Errorf("LightBumperRight has wrong value")
	}
}
