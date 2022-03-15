package gocore

import (
	"fmt"
)

func Cipher_Cezar() {
	source := "This is raw message"
	encrypted_source := ""
	decrypted_source := ""
	rot := 3
	for _, c := range source {
		encrypted_source += string(c + rune(rot))
	}
	fmt.Println("crypted text:", encrypted_source)
	for _, c := range encrypted_source {
		decrypted_source += string(c - rune(rot))
	}
	fmt.Println("decrypted text:", decrypted_source)

}

func Cipher_Vigener() {
	var source string = "This is secret"
	var word string = "key"
	indexOfRune := 0
	var encrypted_source string
	var decrypted_source string
	for i, c := range source {
		if indexOfRune >= len(word) {
			indexOfRune = 0
		}
		encrypted_source += string(c + rune(word[indexOfRune]))
		fmt.Println(i, string(c), indexOfRune, word[indexOfRune], string(c+rune(word[indexOfRune])))
		indexOfRune += 1
	}
	indexOfRune = 0
	for _, c := range encrypted_source {
		if indexOfRune >= len(word) {
			indexOfRune = 0
		}
		decrypted_source += string(c - rune(word[indexOfRune]))
		indexOfRune += 1
	}
	fmt.Println(source)
	fmt.Println(encrypted_source)
	fmt.Println(decrypted_source)
}
