package main

import "net"


func main() {
	net.Dial("tcp",":4411")
	// net.DialTCP() 指定网络只能是tcp，指定本地地址为空随机选择，指定服务端主机地址
	
}