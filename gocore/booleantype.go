package gocore

import (
	"fmt"
	"strings"
)

func BoolContains() {
	fmt.Println("You are in dark cave")
	var command = "Go outside"
	var exit = strings.Contains(command, "outside")
	fmt.Println("You're leaving cave:", exit)
}
