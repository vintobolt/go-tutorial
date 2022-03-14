package gocore

import (
	"fmt"
	"math/big"
)

func BigNumbers() {
	lightSpeed := big.NewInt(299792)
	secondsPerDay := big.NewInt(86400)
	// big int 10^18
	distance := new(big.Int)
	distance.SetString("24000000000000000000", 10)
	fmt.Println(distance)

	seconds := new(big.Int)
	seconds.Div(distance, lightSpeed)

	days := new(big.Int)
	days.Div(seconds, secondsPerDay)
	fmt.Println(days)
}
