package main

import "fmt"


func main() {
	//fmt.Println(isValid("[]()"))
	fmt.Printf("%v",string(byte(92)))
}
func isValid( s string ) bool {
    // write code here
    sbyte := []byte(s)
    if len(sbyte)/2 != 0 || len(sbyte) == 1{
        return false
    }
    stack := make([]byte,0,len(sbyte))
	stack = append(stack, sbyte[0])
    for _,v := range sbyte {
        if v == stack[len(stack)-1] {
            // 出栈
            stack = stack[0:len(stack)-1]
        } else {
            stack = append(stack, v)
        }
    }
    if len(stack) == 0 {
        return true
    }
    return false
}
/**
 * 
 * @param A int整型一维数组 
 * @param B int整型一维数组 
 * @return void
*/
func merge( A []int ,  m int, B []int, n int )  {
    // write code here
	tmp := make([]int,m)
	copy(tmp,A)
	i := 0
	j := 0
	k := 0
    for i < m && j < n {
        if tmp[i] < B[j] {
			A[k] = tmp[i]
			i++
		} else {
			A[k] = B[j]
			j++
		}
		k++
    }
	for ; i < m; i++ {
		A[k] = tmp[i]
		k++
	}
	for ; j < n; j++ {
		A[k] = B[j]
		k++
	}
}