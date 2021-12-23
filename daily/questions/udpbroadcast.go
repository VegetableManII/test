package questions

import (
	"bytes"
	"log"
	"net"
)

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

	data := make([]byte, 512)
	for {
		n, remote, e := l.ReadFromUDP(data)
		if e != nil {
			log.Panicln(e)
		}
		if remote != nil {
			log.Println("client", remote, n)
			log.Println("msg", string(data))
		}
	}
}

var msg = []byte{0x01, 0x01, 0x00, 0x00, 0x12, 0x34, 0x56, 0x78}

// 广播客户端主动发送广播消息
func UdpBroadcastSend() {
	buf := make([]byte, 1024)
	buffer := bytes.NewBuffer(buf)

	la, err := net.ResolveUDPAddr("udp4", "192.168.12.151:12345")
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
	// reader := bufio.NewReader(os.Stdin)
	// text, err := reader.ReadBytes('\n')
	for {
		buffer.Reset()
		_, _, err = conn.ReadFromUDP(buffer.Bytes())
		if err != nil {
			log.Panicln(err)
		}
		log.Println(buffer.String())
		buffer.Reset()
		buffer.Write(msg)
		_, err := conn.Write(buffer.Bytes())
		if err != nil {
			log.Panicln(err)
		}
		//reader.Reset(os.Stdin)

	}
}
