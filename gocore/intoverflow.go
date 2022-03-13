package gocore

import (
	"fmt"
	"time"
)

func OverflowInt() {
	// unsigned int overflow
	var red uint8 = 255
	red++
	fmt.Println(red) // 255 + 1 -> 0

	// Signed int overflow
	var number int8 = 127
	number++
	fmt.Println(number) // 127 + 1 -> -127
	//overflow not throw err, working as circle

	// Date
	future := time.Unix(12622780800, 0)
	fmt.Println(future)
}
