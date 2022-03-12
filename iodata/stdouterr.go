package iodata

import (
	"fmt"
	"io"
	"os"
)

func IOERRImpl() {
	// stdout
	io.WriteString(os.Stdout, "This is string for stdout\n")
	io.WriteString(os.Stderr, "This is string for stderr\n")

	// Stdout/err implementation
	buf := []byte{0xAF, 0xFF, 0xFE}
	for i := 0; i < 200; i++ {
		if _, e := os.Stdout.Write(buf); e != nil {
			panic(e)
		}
	}
	fmt.Fprintln(os.Stdout, "\n")
}
