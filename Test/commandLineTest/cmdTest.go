package main

import (
	"fmt"
	"os"
)

func main() {
	for i, cmd := range os.Args {
		fmt.Printf("arg[%d]:%s\n", i, cmd)
	}
}
