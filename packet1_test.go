package irobot

import "testing"

func TestExtract1(t *testing.T) {
	data := make([]byte, 10)
	data[0] = 0b00000010 // Packet7
	data[1] = 0x01       // Packet8
	data[2] = 0x01       // Packet9
	data[3] = 0x01       // Packet10
	data[4] = 0x00       // Packet11
	data[5] = 0x00       // Packet12
	data[6] = 0x01       // Packet13
	data[7] = 0b00011000 // Packet14
	data[8] = 42         // Packet15
	data[9] = 0x00       // Packet16
	packet := makePacket1().(*Packet1)
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
}
