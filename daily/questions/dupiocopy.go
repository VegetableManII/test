package questions

import (
	"io"
	"log"
	"os"
	"time"
)

func DupIoCopy() {
	src := os.Stdin

	dst1, err := os.OpenFile("/Users/jack/Documents/GitHub/test/daily/questions/test1.log", os.O_WRONLY|os.O_CREATE, os.ModeAppend)
	if err != nil {
		log.Fatalln(err)
	}
	dst2, err := os.OpenFile("/Users/jack/Documents/GitHub/test/daily/questions/test2.log", os.O_WRONLY|os.O_CREATE, os.ModeAppend)
	if err != nil {
		log.Fatalln(err)
	}

	go copytofile(src, dst2, "file2")
	go copytofile(src, dst1, "file1")

	time.Sleep(10 * time.Second)

}

func copytofile(src, dest *os.File, name string) {
	_, err := io.Copy(dest, src)
	if err != nil {
		log.Println(name, err)
	}
}
