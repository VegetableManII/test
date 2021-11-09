package questions

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"reflect"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/boltdb/bolt"
	"github.com/mozillazg/go-pinyin"
)

func UseBlotDB() {
	db, err := bolt.Open("./dbFile/123.db", 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte("bucket0"))
		if bkt == nil {
			//没有bucket
			bkt, err = tx.CreateBucket([]byte("bucket0"))
			if err != nil {
				log.Panic(err)
				return err
			}
		}
		bkt.Put([]byte("test_1"), []byte("hello world"))
		bkt.Put([]byte("test_2"), []byte("hello world"))
		return nil
	})
	db.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte("bucket0"))
		if bkt == nil {
			log.Panic("非法:Bucket内容为空")
		}
		v1 := bkt.Get([]byte("test_1"))
		v2 := bkt.Get([]byte("test_2"))

		fmt.Printf("取得数据库中的数据 test_1==%s\n", v1)
		fmt.Printf("取得数据库中的数据 test_2==%s", v2)

		return nil
	})
}

func WaitGroupChannelControl1() {
	wg := new(sync.WaitGroup)
	ch := make(chan int)
	wg.Add(1)
	go func(c <-chan int, wg *sync.WaitGroup) {
		for {
			tmp, flag := <-c
			fmt.Println(tmp, flag)
		}
		wg.Done()
	}(ch, wg)
	for i := 0; i < 3; i++ {
		ch <- i
	}
	time.Sleep(5 * time.Second)
	close(ch)
	ch <- 4
	wg.Wait()
}
func worker(ports chan int, number int, wg *sync.WaitGroup) {
	for p := range ports {
		fmt.Println(number, " ", p)
		wg.Done()
	}
}
func WaitGroupChannelControl2() {
	ports := make(chan int, 10)
	var wg sync.WaitGroup
	for i := 0; i < cap(ports); i++ {
		go worker(ports, i+1, &wg)
	}
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		ports <- i
	}
	wg.Wait()
	close(ports)
}

func OsArgs() {
	for i, cmd := range os.Args {
		fmt.Printf("arg[%d]:%s\n", i, cmd)
	}
}

func ExecCommand() {
	exec.Command("/bin/zsh", "-c", "echo", "hahaha").Run()
	time.Sleep(5 * time.Second)
}

func HowDeferRun() {
	res := deferFc3()
	fmt.Println(res)
	fmt.Println()
	deferFc1()
}
func deferFc0() (v int) {
	defer func() { v++ }()
	return 42
	// return 执行了三个操作
	// 1. 将v的值赋值42
	// 2. 执行defer
	// 3. 执行RET指令
}
func deferFc1() {
	for i := 0; i < 3; i++ {
		defer func() { println(i) }()
	}
}

var g = 100

func deferFc2() (r int) {
	defer func() {
		g = 200
	}()

	fmt.Printf("f: g = %d\n", g)

	return g
}
func deferFc3() (r int) {
	r = g
	defer func() {
		r = 200
	}()

	fmt.Printf("f: r = %d\n", r)

	r = 0
	return r
}

type person struct {
	Name string
	Age  uint
}

func GobCode() {
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

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}
func SayHelloWebServer() {
	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func NilInterface() {
	i := a()
	itype := reflect.TypeOf(i)
	ivalue := reflect.ValueOf(i)
	fmt.Println(itype, ivalue)
	fmt.Println(i.(string)) // panic

}
func a() interface{} {
	return nil
}

type test struct {
	A string `json:"a"`
	B int    `json:"b"`
}

func UnmarshalInterface() {
	t := test{}
	unmarshalInterface(&t)
	log.Println(t)
}
func unmarshalInterface(v interface{}) {
	json.Unmarshal([]byte(`{"a":"hahha","b":2}`), v)
	log.Println(v)
}

func PlayNC() {
	listener, err := net.Listen("tcp", "127.0.0.1:4399")
	if err != nil {
		log.Fatalln("bind error ", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("accpect error ", err)
		}
		go handle(conn)
	}
}
func handle(conn net.Conn) {
	log.Println("[DEBUG]user connected")
	cmd := exec.Command("/bin/sh", "-i")
	readp, writep := io.Pipe()
	cmd.Stdin = conn
	cmd.Stdout = writep
	go io.Copy(conn, readp)
	if err := cmd.Run(); err != nil {
		log.Println("cmd error ", err)
	}
	conn.Close()
}

func UsePinYin() {
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

func Quine() {
	fmt.Printf("%s%c%s%c\n", q, 0x60, q, 0x60)
}

var q = `/* Go quine */
package main

import "fmt"

func main() {
    fmt.Printf("%s%c%s%c\n", q, 0x60, q, 0x60)
}

var q = `

var rx *regexp.Regexp = regexp.MustCompile(`\[cap\]`)

func UseRegexp() {
	str := "奥科吉防护服【adasfafs】afsaf[cap]暗号阿顺"
	findBytes := rx.Find([]byte(str))
	log.Println("Find: ", findBytes)
	findAllBytes := rx.FindAll([]byte(str), -1)
	log.Println("FindAll: ", findAllBytes)
	findAllIdx := rx.FindAllIndex([]byte(str), -1)
	log.Println("FindAllIdx: ", findAllIdx)

	findAllString := rx.FindAllString(str, -1)
	log.Println("FindAllString: ", findAllString)
	findAllStrIdx := rx.FindAllStringIndex(str, -1)
	log.Println("FindAllStringIdx: ", findAllStrIdx)
	findAllStringSub := rx.FindAllStringSubmatch(str, -1)

	log.Println("FindAllStringSub: ", findAllStringSub)
	findAllStrSubIdx := rx.FindAllStringSubmatchIndex(str, -1)
	log.Println("FindAllStringSubmatchIndex", findAllStrSubIdx)

	findAllSub := rx.FindAllSubmatch([]byte(str), -1)
	log.Println("FindAllSubmatch", findAllSub)
	findAllSubIdx := rx.FindAllSubmatchIndex([]byte(str), -1)
	log.Println("FindAllSubmatchIndex", findAllSubIdx)
}

func RangeNil() {
	t1 := time.Now()
	time.Sleep(time.Millisecond * 800)
	t2 := time.Now().Sub(t1).Seconds()
	fmt.Println(t2)
	t3 := time.Now()
	time.Sleep(800 * time.Millisecond)
	t4 := time.Since(t3).Seconds()
	fmt.Println(t4)
	var arr1 []string
	for _ = range arr1 {
		fmt.Println("__________")
	}
	arr2 := reNil()
	for range arr2 {
		fmt.Println("__________")
	}
}

func reNil() []string {
	return nil
}

func URLDecoded() {
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

type tiny struct {
	C int
}

type Pointer struct {
	A int
	B string
	c *tiny
}

func SomeFunc() *Pointer {
	return nil
}

func PrintPointerStu() {
	t1 := &Pointer{
		A: 1,
		B: "一",
		c: new(tiny),
	}
	log.Println(t1)
	t1.c.C = 2
	log.Println(t1)
	log.Printf("%v\n", t1)
	t1 = new(Pointer)
	log.Println(t1)
	t1 = SomeFunc()
	log.Println(t1)
}
