package keyboard

import "fmt"

func Keyboard() {
	var name string
	fmt.Println("Whats your name?")
	fmt.Scanf("%s\n", &name) // save string into variable

	var age int
	fmt.Println("How old are you?")
	fmt.Scanf("%d\n", &age) // save integer into variable

	fmt.Printf("Hi, %s, your age - %d\n", name, age)
}
