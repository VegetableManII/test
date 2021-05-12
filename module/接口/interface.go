package main

import (
	"fmt"
	"reflect"
	"time"
)

type A interface {
	Go()
}
type B struct {
	fied1 int
	fied2 string
	fied  map[int]string
}

func (b *B) Go() {
}

var _ A = (*B)(nil)

func wr(ch chan<- int) {
	time.Sleep(time.Second * 2)
	ch <- 1
	time.Sleep(time.Second)
	close(ch)
}

func main() {
	test(&B{})
}

func test(v interface{}) {
	a := reflect.ValueOf(v)
	a1 := a.Elem()
	b := reflect.TypeOf(v)
	b1 := b.Elem()
	for i := 0; i < a1.NumField(); i++ {
		fmt.Println(a1.Field(i))
	}
	fmt.Println()
	for i := 0; i < b1.NumField(); i++ {
		fmt.Println(b1.Field(i))
	}
}
