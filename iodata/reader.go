package iodata

import (
	"fmt"
	"os"
)

func ReadFromStdin() {
	for {
		data := make([]byte, 8)
		n, err := os.Stdin.Read(data)
		if err == nil && n > 0 {
			process(data)
		} else {
			break
		}
	}
}

func process(data []byte) {
	fmt.Printf("Recieved: %X	%s\n", data, string(data))
}
