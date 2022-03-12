package iodata

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func OpenFile() {
	path_read := "/Users/vintobolt/Projects/golang/go-tutorial/README.md"
	path_create := "/Users/vintobolt/Projects/golang/go-tutorial/test.txt"

	// Read file
	f, err := os.Open(path_read)
	if err != nil {
		panic(err)
	}

	c, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("### File content ###\n%s\n", string(c))
	f.Close()

	f, err = os.OpenFile(path_create, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	io.WriteString(f, "Test string")
	f.Close()
}
