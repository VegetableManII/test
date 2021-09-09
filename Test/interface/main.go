package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := a()
	itype := reflect.TypeOf(i)
	ivalue := reflect.ValueOf(i)
	fmt.Println(itype, ivalue)
	fmt.Println(i.(string)) // panic

}

func a() interface{} {
	return nil
}
