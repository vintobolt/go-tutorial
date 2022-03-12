package iodata

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func Multiwr() {
	buf := bytes.NewBuffer([]byte{})
	f, err := os.OpenFile("sample.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	wr := io.MultiWriter(buf, f)
	_, err = io.WriteString(wr, "Hi, go is nice!")
	if err != nil {
		panic(err)
	}
	fmt.Println("buffer contain: " + buf.String())
}
