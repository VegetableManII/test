package main

import (
	"io"
	"net"
	"sync"
)

var wg sync.WaitGroup

func main() {
	listener, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		return
	}
	wg.Add(1)
	// 开启线程监听8001的代理请求
	// go func() 为go语言语法，开启一个线程运行func()
	go handle(listener)
	wg.Wait()
	return
}
func handle(listener net.Listener) {
	for {
		var handlewg sync.WaitGroup
		// 拿到8001的连接代理的请求
		connection, err := listener.Accept()
		if err == nil {
			// 开启线程创建到目的服务器的连接
			go func() {
				// 向目的8002发起连接建立的请求
				remote, err := net.Dial("tcp", "localhhost:8002")
				if err != nil {
					return
				}
				// 连接建立完成后，拷贝双方的io以实现隧道通信
				handlewg.Add(1)
				go copy(remote, connection, &handlewg)
				handlewg.Add(1)
				go copy(connection, remote, &handlewg)
				remote.Close()
				connection.Close()
			}()
		}
	}
}
func copy(from, to net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	if _, err := io.Copy(to, from); err != nil {
		// io 复制出错，返回
		return
	}
}
