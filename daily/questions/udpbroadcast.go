package questions

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
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

type EpcMsg struct {
	_type   byte
	_method byte
	_size   uint16
	data    [1020]byte
}

// 接收广播消息
func UdpBroadcastAsyncReceive() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// local address
	la, err := net.ResolveUDPAddr("udp4", "0.0.0.0:65532")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, e := net.ListenUDP("udp4", la)
	if e != nil {
		log.Panicln(e)
	}
	log.Println("listen on", conn.LocalAddr().String())
	data := make([]byte, 1024)
	n, remote, e := conn.ReadFromUDP(data)
	if e != nil {
		log.Panicln(e)
	}
	log.Printf("R[%v]: %v(string)", n, string(data[:n]))
	n, e = conn.WriteToUDP([]byte{0x0F, 0x0F, 0x0F, 0x0F}, remote)
	myid := make([]byte, 4)
	n, remote, e = conn.ReadFromUDP(myid)
	if e != nil {
		log.Panicln(e)
	}
	log.Printf("R[%v]: %v(string)", n, string(myid[:n]))
	log.Println(myid)
	go func() {
		for {
			n, remote, e = conn.ReadFromUDP(data)
			if e != nil {
				log.Panicln(e)
			}
			log.Printf("R[%v]: %v(string)", n, string(data[:n]))
			time.Sleep(3 * time.Second)
		}
	}()
	registWithAuth := []byte("REGISTER sip:hebeiyidong.3gpp.net SIP/2.0\r\n" +
		"Via: SIP/2.0/UDP 10.255.1.111:5090;branch=z9hG4bK199912928954841999\r\n" +
		`From: "jiqimao" <sip:jiqimao@hebeiyidong.3gpp.net>;tag=690713` + "\r\n" +
		`To: "jiqimao" <sip:jiqimao@hebeiyidong.3gpp.net>;tag=690711` + "\r\n" +
		"Call-ID: RgeX-136783086082016@10.255.1.111\r\n" +
		"CSeq: 3 REGISTER\r\n" +
		"Contact: <sip:jiqimao@10.255.1.111:5090>\r\n" +
		// `Authorization: Digest username="ng40user11"` + "\r\n" + 有P-CSCF产生
		"Allow: INVITE,ACK,OPTIONS,CANCEL,BYE,PRACK,UPDATE,SUBSCRIBE,NOTIFY\r\n" +
		"Max-Forwards: 70\r\n" +
		"Expires: 600000\r\n" +
		"Content-Length: 0\r\n" +
		"\r\n")

	da := append(myid, registWithAuth...)
	n, e = conn.WriteToUDP(da, remote)
	if e != nil {
		log.Panicln(e)
	}
	log.Printf("S[%v]: %v\n", n, string(da))
	<-quit
}

func stringToBytes(s string) []byte {
	b := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		b = append(b, s[i]-0x30)
	}
	return b
}

var msg = []byte{0x01, 0x00, 0x00, 0x00, 0x12, 0x34, 0x56, 0x78}

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
