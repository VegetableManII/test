package main

import (
	"log"
	"strings"
)

func main() {
	str1 := "123"
	str3 := "123,456,789"
	items1 := strings.Split(str1, ",")
	items3 := strings.Split(str3, ",")
	mmap1 := make(map[string]struct{}, 1)
	for _, v := range items1 {
		mmap1[v] = struct{}{}
	}
	mmap3 := make(map[string]struct{}, 1)
	for _, v := range items3 {
		mmap3[v] = struct{}{}
	}
	if v, ok := mmap1["123"]; ok {
		log.Println("123存在", v, mmap1)
	} else {
		log.Println("123不存在", mmap1)
	}
	if v, ok := mmap3["123"]; ok {
		log.Println("123存在", v, mmap3)
	} else {
		log.Println("123不存在", mmap3)
	}
}
