package main

import (
	"fmt"
	"os"

	"github.com/web-assets/go-jsmin"
)

func main() {
	for i, arg := range os.Args {
		if i != 0 {
			fmt.Printf("// %s\n", arg)
		}
	}

	jsmin.Min(os.Stdin, os.Stdout)
}
