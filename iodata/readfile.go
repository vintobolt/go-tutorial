package iodata

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile() {
	path := "/Users/vintobolt/Projects/golang/go-tutorial/test.txt"

	fmt.Println("### Read as reader ###")
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Reading file with reader
	wr := bytes.Buffer{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		wr.WriteString(sc.Text())
	}
	fmt.Println(wr.String())

	fmt.Println("### ReadFile into variable ###")
	fContent, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fContent))
}
