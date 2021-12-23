package questions

import (
	"fmt"
	"log"
	"net"
	"time"
)

// '0' = 48 = 0x30

// 广播服务端接收消息
func UdpBroadcastReceive() {
	l, e := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 12345,
	})
	if e != nil {
		log.Panicln(e)
	}
	log.Println("listen on", l.LocalAddr().String())

	data := make([]byte, 1024)
	for {
		n, remote, e := l.ReadFromUDP(data)
		if e != nil {
			log.Panicln(e)
		}
		if remote != nil {
			log.Println("remote", remote, n)
			log.Println("msg", string(data))
			text := ""
			fmt.Scanln(&text)
			if e != nil {
				log.Panicln(e)
			}
			log.Println(text)
			_, e = l.WriteToUDP(stringToBytes(text), remote)
			if e != nil {
				log.Panicln(e)
			}
		} else {
			time.Sleep(5 * time.Second)
		}
		data = data[:0] // 清空缓冲区
	}
}

var msg = []byte{0x01, 0x01, 0x00, 0x00, 0x12, 0x34, 0x56, 0x78}

func stringToBytes(s string) []byte {
	b := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		b = append(b, s[i]-0x30)
	}
	return b
}

// 广播客户端主动发送广播消息
func UdpBroadcastSend() {
	buf := make([]byte, 0, 1024)
	la, err := net.ResolveUDPAddr("udp4", ":11111")
	if err != nil {
		log.Panicln(err)
	}
	ra, err := net.ResolveUDPAddr("udp4", "255.255.255.255:12345")
	if err != nil {
		log.Panicln(err)
	}
	conn, err := net.DialUDP("udp4", la, ra)
	if err != nil {
		log.Panicln(err)
	}

	conn.Write([]byte("start work"))
	for {
		_, _, err = conn.ReadFromUDP(buf)
		if err != nil {
			log.Panicln(err)
		}
		log.Println(string(buf))
		_, err := conn.Write(msg)
		if err != nil {
			log.Panicln(err)
		}
		buf = buf[:0]
	}
}
