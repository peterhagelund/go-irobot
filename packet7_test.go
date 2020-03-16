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

func TestExtract7(t *testing.T) {
	packet := makePacket7().(*Packet7)
	data := make([]byte, 1)
	data[0] = 0b00000000
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.BumpRight != false {
		t.Errorf("BumpRight has wrong value")
	}
	if packet.BumpLeft != false {
		t.Errorf("BumpLeft has wrong value")
	}
	if packet.WheelDropRight != false {
		t.Errorf("WheelDropRight has wrong value")
	}
	if packet.WheelDropLeft != false {
		t.Errorf("WheelDropLeft has wrong value")
	}
	data[0] = 0b00001111
	if err := packet.Extract(data, 0); err != nil {
		t.Error(err)
	}
	if packet.BumpRight != true {
		t.Errorf("BumpRight has wrong value")
	}
	if packet.BumpLeft != true {
		t.Errorf("BumpLeft has wrong value")
	}
	if packet.WheelDropRight != true {
		t.Errorf("WheelDropRight has wrong value")
	}
	if packet.WheelDropLeft != true {
		t.Errorf("WheelDropLeft has wrong value")
	}
}
