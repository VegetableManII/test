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

type DhSanTongReport struct {
	Msgid      string `json:"msgid"`
	Phone      string `json:"phone"`
	Status     string `json:"status"`
	Desc       string `json:"desc"`
	Wgcode     string `json:"wgcode"`
	SubmitTime string `json:"submitTime"`
	SendTime   string `json:"sendTime"`
	Time       string `json:"time"`
	SmsCount   int    `json:"smsCount"`
	SmsIndex   int    `json:"smsIndex"`
}

type DhSanTongStatusSt struct {
	Result  string             `json:"result"`
	Desc    string             `json:"desc"`
	Reports []*DhSanTongReport `json:"reports"`
}

func main() {
	str := `{
		"result":"0",
		"desc":"成功",
		"reports":[{
			"msgid":"2c92825934837c4d0134837dcba00150",
			"phone":"13534698345",
			"status":"0",
			"desc":"成功",
			"wgcode":"DELIVRD",
			"submitTime":"2015-03-17 16:32:16",
			"sendTime":"2015-03-17 16:32:17",
			"time":"2015-03-17 16:32:20",
			"smsCount":1,
			"smsIndex":1
			},{
			"msgid":"2c92825934837c4d0134837dcba02884d",
			"phone":"13917599647",
			"status":"0",
			"desc":"成功",
			"wgcode":"DELIVRD",
			"submitTime":"2015-03-17 16:32:20",
			"sendTime":"2015-03-17 16:32:20",
			"time":"2015-03-17 16:32:21",
			"smsCount":1,
			"smsIndex":1
			}]
	}`

	callbackSt := new(DhSanTongStatusSt)
	callbackSt.Reports = make([]*DhSanTongReport, 0, 1)
	if e := json.Unmarshal([]byte(str), &callbackSt); e != nil {
		fmt.Println("json解析失败", e)
	}
	fmt.Println(callbackSt)
	// tmp := Rr{
	// 	C: "哈哈",
	// }
	// jsbytes, err := json.Marshal(tmp)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// r1 := Rr{}
	// r1ptr := new(Rr)
	// if err = json.Unmarshal(jsbytes, r1); err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("r1 %v\n", r1)
	// if err = json.Unmarshal(jsbytes, &r1); err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("&r1 %v\n", r1)
	// if err = json.Unmarshal(jsbytes, r1ptr); err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("r1ptr %v\n", r1)
	// if err = json.Unmarshal(jsbytes, &r1ptr); err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("&r1ptr %v\n", r1)

	// t1 := Tt{}
	// t1ptr := new(Tt)
	// if err = json.Unmarshal(jsbytes, &t1); err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("&t1 %v\n", t1)
	// if err = json.Unmarshal(jsbytes, t1ptr); err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("t1ptr %v\n", t1)
	// if err = json.Unmarshal(jsbytes, &t1ptr); err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("&t1ptr %v\n", t1)

}
