package main

import (
	"fmt"
	"strings"

	"github.com/mozillazg/go-pinyin"
)

func main() {
	hans := "abc招商银行"
	tunes := pinyin.LazyPinyin(hans, pinyin.Args{
		Style:     pinyin.Normal,
		Heteronym: false,
		Separator: "",
		Fallback: func(r rune, a pinyin.Args) []string {
			return []string{string(r)}
		},
	})
	res := strings.Join(tunes, "")
	fmt.Println(fmt.Sprintf("captcha_%s", res))
}
