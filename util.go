package irobot

func bytesToUint16(hi, lo byte) uint16 {
	return uint16(hi)<<8 | uint16(lo)
}

func uint16ToBytes(value uint16) (byte, byte) {
	return byte(value >> 8), byte(value & 0x00ff)
}

func bytesToInt16(hi, lo byte) int16 {
	return int16(hi)<<8 | int16(lo)
}

func int16ToBytes(value int) (byte, byte) {
	return byte(value >> 8), byte(value & 0x00ff)
}
