package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scanln(&str)
	var sub string
	fmt.Scanln(&sub)
	fmt.Println(strings.Index(str, sub))
}
