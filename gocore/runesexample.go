package gocore

import "fmt"

// "runes" alias for int32.

func Runes() {
	emoji := []rune("привет😀")

	for i := 0; i < len(emoji); i++ {
		fmt.Println(emoji[i], string(emoji[i]))
	}
	fmt.Println("Len of runes:", len(emoji))
	fmt.Println("Len of string:", len(string(emoji)))
}
