package gocore

import "fmt"

func IfElseUsage() {
	var command = "go East"

	if command == "go East" {
		fmt.Println("You're going to mountain.")
	} else if command == "come in" {
		fmt.Println("You're comming in cave where will live.")
	} else {
		fmt.Println("While nothing is clear")
	}
}

// || even is one of confition is true
// && both conditions must be true
func LogicOperatorsUsage() {
	fmt.Println("Now 2100 year. It is leap.")

	var year = 2100
	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)

	if leap {
		fmt.Println("This year is leap.")
	} else {
		fmt.Println("Unfortunately no, this year is not leap.")
	}
}

func SwitchUsageWithVar() {
	fmt.Println("This is entry to cave and path to East")
	var command = "Come in"

	switch command {
	case "Come into cave", "Come in":
		fmt.Println("You're in dimly lit cave")
	case "read the sign":
		fmt.Println("The sign says 'Minors are not allowed to enter'.")
	default:
		fmt.Println("Nothing is clear")
	}
}

func SwitchUsageWithoutVar() {
	var room = "lake"

	switch {
	case room == "cave":
		fmt.Println("You're in dimly lip cave")
	case room == "lake":
		fmt.Println("The ice seems to be strong enough")
		fallthrough // goto next case
	case room == "depth":
		fmt.Println("Water is so cool")
	}
}
