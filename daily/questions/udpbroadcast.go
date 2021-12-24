package questions

import (
	"fmt"
	"log"
	"net"
	"time"
)

// '0' = 48 = 0x30

// 接收广播消息
func UdpBroadcastReceive() {
	// local address
	la, err := net.ResolveUDPAddr("udp4", "0.0.0.0:12345")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, e := net.ListenUDP("udp4", la)
	if e != nil {
		log.Panicln(e)
	}
	log.Println("listen on", conn.LocalAddr().String())
	// data := make([]byte, 0, 64) // ReadFromUDP always return n = 0
	data := make([]byte, 32)
	text := ""
	for {
		n, remote, e := conn.ReadFromUDP(data)
		if e != nil {
			log.Panicln(e)
		}
		log.Printf("R[%v]: %v\n", n, data[:n])
		// io.Copy(os.Stdout, conn)
		_, e = fmt.Scanln(&text)
		if e != nil {
			log.Panicln(e)
		}
		bs := stringToBytes(text)
		n, e = conn.WriteToUDP(bs, remote)
		if e != nil {
			log.Panicln(e)
		}
		log.Printf("S[%v]: %v\n", n, bs)
		data = data[:0]
	}
}

// 接收广播消息
func UdpBroadcastAsyncReceive() {
	// local address
	la, err := net.ResolveUDPAddr("udp4", "0.0.0.0:12345")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, e := net.ListenUDP("udp4", la)
	if e != nil {
		log.Panicln(e)
	}
	log.Println("listen on", conn.LocalAddr().String())

	// data := make([]byte, 0, 64) // ReadFromUDP always return n = 0
	data := make([]byte, 16)
	n, remote, e := conn.ReadFromUDP(data)
	if e != nil {
		log.Panicln(e)
	}
	log.Printf("R[%v]: %v\n", n, data[:n])
	go func() {
		for {
			n, remote, e = conn.ReadFromUDP(data)
			if e != nil {
				log.Panicln(e)
			}
			log.Printf("R[%v]: %v\n", n, data[:n])
			time.Sleep(500 * time.Millisecond)
		}
	}()
	for {
		n, e = conn.WriteToUDP(msg, remote)
		if e != nil {
			log.Panicln(e)
		}
		log.Printf("S[%v]: %v\n", n, msg)
		time.Sleep(time.Second)
	}
}

func stringToBytes(s string) []byte {
	b := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		b = append(b, s[i]-0x30)
	}
	return b
}

var msg = []byte{0x01, 0x01, 0x00, 0x00, 0x12, 0x34, 0x56, 0x78}

// 广播客户端主动发送广播消息
func UdpBroadcastSend() {
	buf := make([]byte, 32)
	la, err := net.ResolveUDPAddr("udp4", ":12345")
	if err != nil {
		log.Panicln(err)
	}
	ra, err := net.ResolveUDPAddr("udp4", "255.255.255.255:11111")
	if err != nil {
		log.Panicln(err)
	}
	conn, err := net.ListenUDP("udp4", la)
	if err != nil {
		log.Panicln(err)
	}

	conn.WriteToUDP([]byte("start work"), ra)
	for {
		n, r, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Panicln(err)
		}
		if r == nil || n == 0 {
			log.Println("empty message")
			time.Sleep(2 * time.Second)
		}
		log.Println(buf)
		_, err = conn.WriteToUDP(buf, ra)
		if err != nil {
			log.Panicln(err)
		}
		buf = buf[:0]
	}
}
