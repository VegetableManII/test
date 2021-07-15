package main

import "fmt"

func main() {
	fmt.Printf("%x\n", -1)
	var testNum int64 = -1 << 63
	fmt.Printf("%x\n", testNum)
	testNum = testNum ^ -1
	fmt.Printf("%x\n", testNum)
	testNum = testNum & -1
	fmt.Printf("%x\n", testNum)
	testNum = testNum | -1
	fmt.Printf("%x\n", testNum)
}
