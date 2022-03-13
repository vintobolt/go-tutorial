package gocore

import "fmt"

func Hexa() {
	var red, green, blue uint8 = 0x00, 0x8d, 0xd5
	fmt.Printf("%02x %02x %02x\n", red, green, blue)
	// 02x | 0 - zero indent, 2 count of signs
}
