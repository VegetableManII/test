package main

import "fmt"

func main() {
	if true {
		defer fmt.Printf("1")
	} else {
		defer fmt.Printf("2")
	}
	fmt.Printf("3")
}
