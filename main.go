package main

import (
	"fmt"
	"os"
)

func main() {
	for i, v := range os.Args {
		if i > 0 {
			fmt.Printf("%d: %s\n", i, v)
		}
	}
}
