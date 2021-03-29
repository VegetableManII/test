package meituan

import (
	"fmt"
	"testing"
)
type t struct {
	index *[]int
}

func TestRemeberNumber(t *testing.T) {
	remeberNumber()
}
// 第一行输入需要记忆的数字的数量和要提问的次数
// 第二行输入所有的数字
// 第三行输入提问的数字
// 输出
// 输出提问的数字出现的位置
func remeberNumber() {
	var n, q int
	fmt.Scanf("%d %d\n",&n,&q)
	num := make(map[int]*t,0)
	for i := 0;i<n ;i++  {
		var tmp int
		fmt.Scanf("%d",&tmp)
		if _,exist := num[tmp]; exist {
			*(num[tmp].index) = append((*num[tmp].index),i+1)
		} else {
			s := make([]int,0)
			s = append(s, i+1)
			number := &t{&s}
			num[tmp] = number
		}
	}
	question := make([]int,0,q)
	for i :=0;i<q ;i++  {
		fmt.Scanln(question[i])
	}
	for v := range question{
		for value := range *(num[v].index) {
			fmt.Printf("%d ",value)
		}
		fmt.Println()
	}

}
// 第一行 输入有几组数据
// 第二行 输入这组数据中有几个字符
// 第三行 输入这组数据的实际内容
// 输入 如果是接近回文则输出替换一个字符后的回文串
//      如果无法转换成回文则输出替换一个字符后的最小值
func findBestNumber() {
	var t, n int
	fmt.Scanln(&t)
	for i := 0; i < t; i++ {
		fmt.Scanln(&n)
		number := make([]byte, 0, n)
		var tmp byte
		for i := 0; i < n; i++ {
			fmt.Scanf("%c",&tmp)
			if tmp == '\n' {
				break
			}
			number = append(number, tmp)
		}
		fmt.Println(len(number))
		fmt.Printf("%s\n",number)
		ok := huiwen(number)
		if !ok {
			zuixiao(number)
		}
		fmt.Println(len(number))
		fmt.Printf("%+v\n",number)
		fmt.Printf("%s\n",number)
	}
}

func huiwen(num []byte) bool {
	if len(num) == 1 {
		return true
	}
	i, j, count, index := 0, len(num)-1, 0, -1
	for i <= j {
		if num[i] != num[j] {
			count ++
			index = i
			if count > 1 {
				return  false
			}
		}
		i++
		j--
	}
	if index != -1 {
		num[index] = num[len(num) - index -1]
	}
	return true
}
func zuixiao(num []byte) {
	for i := range num {
		if num[i] != '0' {
			num[i] = '0'
			return
		}
	}
}
