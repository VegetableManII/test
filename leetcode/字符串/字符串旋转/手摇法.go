package main

import (
	"fmt"
	"math"
)

func ReverrseString(s []byte, from, to int) []byte {
	sbyte := []byte(s)
	for from < to {
		t := sbyte[from]
		sbyte[from] = s[to]
		sbyte[to] = t
		from++
		to--
	}
	return sbyte
}

/*
将一个字符串分成X和Y两个部分，
在每部分字符串上定义反转操作，如X^T，即把X的所有字符反转
（如，X=”abc”，那么X^T=”cba”），
那么就得到下面的结论：(X^TY^T)^T=YX，
显然就解决了字符串的反转问题。
*/

// LeftRotateString翻转字符串 m 右移位数
func LeftRotateString(s string, n, m int) []byte {
	m = m % n // 当m移动的位数超过n通过取余得到同样的效果
	return ReverrseString(ReverrseString(ReverrseString([]byte(s), 0, m-1), m, n-1), 0, n-1)
	//反转[0..m - 1]，套用到上面举的例子中，就是X->X^T，即 abc->cba
	//反转[m..n - 1]，例如Y->Y^T，即 def->fed
	//反转[0..n - 1]，即如整个反转，(X^TY^T)^T=YX，即 cbafed->defabc。
}
func main() {
	/* s := "helloworld!"
	_ = LeftRotateString(s, len(s), 1) */
	fmt.Println(math.MaxInt32)
}
