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

func TestExtract18(t *testing.T) {
	packet := makePacket18().(*Packet18)
	data := make([]byte, 1)
	data[0] = 0b00000000
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.Clean != false {
		t.Errorf("Clean has wrong value")
	}
	if packet.Spot != false {
		t.Errorf("Spot has wrong value")
	}
	if packet.Dock != false {
		t.Errorf("Dock has wrong value")
	}
	if packet.Minute != false {
		t.Errorf("Minute has wrong value")
	}
	if packet.Hour != false {
		t.Errorf("Hour has wrong value")
	}
	if packet.Day != false {
		t.Errorf("Day has wrong value")
	}
	if packet.Schedule != false {
		t.Errorf("Schedule has wrong value")
	}
	if packet.Clock != false {
		t.Errorf("Clock has wrong value")
	}
	data[0] = 0b11111111
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.Clean != true {
		t.Errorf("Clean has wrong value")
	}
	if packet.Spot != true {
		t.Errorf("Spot has wrong value")
	}
	if packet.Dock != true {
		t.Errorf("Dock has wrong value")
	}
	if packet.Minute != true {
		t.Errorf("Minute has wrong value")
	}
	if packet.Hour != true {
		t.Errorf("Hour has wrong value")
	}
	if packet.Day != true {
		t.Errorf("Day has wrong value")
	}
	if packet.Schedule != true {
		t.Errorf("Schedule has wrong value")
	}
	if packet.Clock != true {
		t.Errorf("Clock has wrong value")
	}
	data[0] = 0b01010101
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.Clean != true {
		t.Errorf("Clean has wrong value")
	}
	if packet.Spot != false {
		t.Errorf("Spot has wrong value")
	}
	if packet.Dock != true {
		t.Errorf("Dock has wrong value")
	}
	if packet.Minute != false {
		t.Errorf("Minute has wrong value")
	}
	if packet.Hour != true {
		t.Errorf("Hour has wrong value")
	}
	if packet.Day != false {
		t.Errorf("Day has wrong value")
	}
	if packet.Schedule != true {
		t.Errorf("Schedule has wrong value")
	}
	if packet.Clock != false {
		t.Errorf("Clock has wrong value")
	}
}
