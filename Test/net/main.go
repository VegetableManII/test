package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func process(cnn net.Conn) {
	msg := make([]byte, 256)
	for {
		n, err := cnn.Read(msg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(msg[0:n]))
	}
}
func _main() {
	p := os.Args[1]
	fmt.Println("port" + p)
	listen, err := net.Listen("tcp", "localhost:"+p)
	if err != nil {
		log.Fatal(err)
	}
	for {
		con, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go process(con)
	}
}
