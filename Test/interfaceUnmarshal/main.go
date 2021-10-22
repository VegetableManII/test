package main

import (
	"encoding/json"
	"log"
)

type test struct {
	A string `json:"a"`
	B int    `json:"b"`
}

func main() {
	t := test{}
	unmarshalInterface(&t)
	log.Println(t)
}

func unmarshalInterface(v interface{}) {

	json.Unmarshal([]byte(`{"a":"hahha","b":2}`), v)
	log.Println(v)
}
