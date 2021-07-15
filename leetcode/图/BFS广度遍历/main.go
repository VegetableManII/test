package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("2"))

}

var mmap = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	return bfs(digits)

}
func bfs(digits string) []string {
	// 从队列中取出每一层中的数据依次拼接后再放入队列
	results := make([]string, 0, 0)
	queue := make([]string, 1, len(digits))
	queue[0] = ""
	// 遍历的层次,让每一层对应的 字母串 都和队列中已有字符串进行拼接
	for i := range digits {
		digit := int(digits[i]) - 48
		letters := mmap[digit]
		// 遍历队列,取得队列中已完成拼接的字符串
		for range queue {
			tmpStr := queue[0]
			queue = queue[1:]
			// 遍历这一层的 字母串 将其拼接到队列中的字符串的每一个的尾部
			for j := range letters {
				queue = append(queue, tmpStr+string(letters[j]))
			}
		}
	}
	for len(queue) > 0 {
		results = append(results, queue[0])
		queue = queue[1:]
	}
	return results
}
