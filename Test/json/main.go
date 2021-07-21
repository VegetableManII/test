package main

import (
	"encoding/json"
	"fmt"
)

type Tt struct {
	A string `json:"a"`
	B string `json:"b"`
}

type Rr struct {
	C string `json:"c"`
}

func main() {
	tmp := Rr{
		C: "哈哈",
	}
	jsbytes, err := json.Marshal(tmp)
	if err != nil {
		fmt.Println(err)
	}
	r1 := Rr{}
	r1ptr := new(Rr)
	if err = json.Unmarshal(jsbytes, r1); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("r1 %v\n", r1)
	if err = json.Unmarshal(jsbytes, &r1); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("&r1 %v\n", r1)
	if err = json.Unmarshal(jsbytes, r1ptr); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("r1ptr %v\n", r1)
	if err = json.Unmarshal(jsbytes, &r1ptr); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("&r1ptr %v\n", r1)

	t1 := Tt{}
	t1ptr := new(Tt)
	if err = json.Unmarshal(jsbytes, &t1); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("&t1 %v\n", t1)
	if err = json.Unmarshal(jsbytes, t1ptr); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("t1ptr %v\n", t1)
	if err = json.Unmarshal(jsbytes, &t1ptr); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("&t1ptr %v\n", t1)

}
