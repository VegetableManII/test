package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("2"))
}

var mmap []string = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
var results []string = []string{}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	backtrace(digits, 0, "")
	return results

}
func backtrace(digits string, index int, res string) {
	if index == len(digits) {
		results = append(results, res)
	} else {
		digit := int((digits[index])) - 48
		letters := mmap[digit]
		for i := range letters {
			backtrace(digits, index+1, res+string(letters[i]))
		}
	}

}
