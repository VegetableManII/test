package main

import (
	"fmt"
	"net/url"
)

func main() {
	str := `report=%7B%22result%22%3A%220%22%2C%22reports%22%3A%5B%7B%22submitTime%22%3A%222021-07-21+17%3A10%3A20%22%2C%22phone%22%3A%2215076329742%22%2C%22smsCount%22%3A1%2C%22wgcode%22%3A%22DELIVRD%22%2C%22msgid%22%3A%220bccyp90gilagw_2021072117%22%2C%22smsIndex%22%3A1%2C%22time%22%3A%222021-07-21+17%3A10%3A24%22%2C%22sendTime%22%3A%222021-07-21+17%3A10%3A21%22%2C%22desc%22%3A%22%E5%8F%91%E9%80%81%E6%88%90%E5%8A%9F%22%2C%22status%22%3A%220%22%7D%5D%2C%22desc%22%3A%22%E6%88%90%E5%8A%9F%22%7D`
	if res, e := url.QueryUnescape(str); e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(res)
	}
	// if values, e := url.ParseQuery(str); e != nil {
	// 	fmt.Println(e)
	// } else {
	// 	for _, v := range values {
	// 		if res, e := url.QueryUnescape(v); e != nil {
	// 			fmt.Println(e)
	// 		} else {
	// 			fmt.Println(res)
	// 		}
	// 	}
	// }

}
