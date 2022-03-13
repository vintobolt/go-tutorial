package gocore

import "fmt"

func IntegerBits() {
	var green uint8 = 3
	fmt.Printf("%08b\n", green)
	green++
	fmt.Printf("%08b\n", green)
}
