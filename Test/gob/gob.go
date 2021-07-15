package gobtest

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type person struct {
	Name string
	Age  uint
}

func main() {
	var xiaoMing person
	xiaoMing.Name = "xiaoming"
	xiaoMing.Age = 20

	//编码的数据放在buffer
	var buffer bytes.Buffer

	//使用gob进行序列化
	//定义编码器
	//使用编器及进行编码
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&xiaoMing)
	if err != nil {
		log.Panic("编码出错")
	}
	fmt.Printf("编码后的数据:%v\n", buffer.Bytes())

	//使用gob进行反序列化
	//定义一个解码器
	decoder := gob.NewDecoder(bytes.NewReader(buffer.Bytes()))

	var dm person
	//使用解码器解码
	err = decoder.Decode(&dm)
	if err != nil {
		log.Panic("解码出错")
	}
	fmt.Printf("解码后的数据:%v\n", dm)
}
