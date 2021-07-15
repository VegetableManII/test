package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func _main() {
	str := "[123]文本内容"
	prefix := "[123]"
	fmt.Printf("str: %+v addr: %+v len: %+v\n", str, (*reflect.StringHeader)(unsafe.Pointer(&str)).Data,
		(*reflect.StringHeader)(unsafe.Pointer(&str)).Len)
	str = strings.TrimPrefix(str, prefix)
	fmt.Printf("str: %+v addr: %+v len: %+v\n", str, (*reflect.StringHeader)(unsafe.Pointer(&str)).Data,
		(*reflect.StringHeader)(unsafe.Pointer(&str)).Len)
}
