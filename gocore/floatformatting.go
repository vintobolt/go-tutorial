package gocore

import (
	"fmt"
	"math"
)

func FloatFormat() {
	third := 1.0 / 3
	fmt.Println(third)           // full
	fmt.Printf("%v\n", third)    // too full
	fmt.Printf("%f\n", third)    // 0.333333
	fmt.Printf("%.3f\n", third)  // 0.333
	fmt.Printf("%4.2f\n", third) // 0.33
	// %4 - width, .2 - precision

	thirdF := 1.0 / 3.0
	fmt.Println(thirdF * 3)

	piggyBank := 0.1
	piggyBank += 0.2
	fmt.Println(piggyBank)

	celsius := 21.0
	// Article said that multiplication takes precedence over division
	// But it fixed in latest compilators
	fmt.Print((celsius/5.0*9.0)+32, "° F\n")
	fmt.Print((9.0/5.0*celsius)+32, "° F\n")

	fahrenheit := (celsius * 9.0 / 5.0) + 32.0
	fmt.Print(fahrenheit, "° F\n") // Выводит: 69.8° F

	// Comparing float digits
	fmt.Println("Comparing float digits")
	fmt.Println(piggyBank)
	fmt.Println(math.Abs(piggyBank-0.3) < 0.0001)
}
