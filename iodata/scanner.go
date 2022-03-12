package iodata

import (
	"bufio"
	"fmt"
	"os"
)

func UseScanner() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		txt := sc.Text()
		fmt.Printf("Echo: %s\n", txt)
	}
}
