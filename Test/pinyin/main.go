package main

import (
	"fmt"
	"strings"

	"github.com/mozillazg/go-pinyin"
)

func main() {
	hans := "招商银行"
	tunes := pinyin.LazyPinyin(hans, pinyin.NewArgs())
	res := strings.Join(tunes, "")
	fmt.Println(fmt.Sprintf("captcha_%s", res))
}
