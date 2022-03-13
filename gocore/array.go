package gocore

import "fmt"

func SimpleArr() {
	var x [5]int // array of integer with length 5
	x[4] = 100
	fmt.Println(x)
}

func SummInArr() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))
}

func GreaterSumm() {
	x := [5]float64{98, 93, 77, 82, 83}

	var total float64 = 0
	for _, value := range x { // _ is index
		total += value
	}
	fmt.Println(total / float64(len(x)))
}
