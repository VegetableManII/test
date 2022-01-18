package questions

import "fmt"

func StringStop() {
	b := []byte{53, 54, 55, 56, 57, 59, 0, 0, 0, 0, 0, 0, 0}
	s := string(b)
	fmt.Println("len []byte", len(b))
	fmt.Println("len string", len(s), s)
}
