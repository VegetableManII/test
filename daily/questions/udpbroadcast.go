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

type EpcMsg struct {
	_type   byte
	_method byte
	_size   uint16
	data    [1020]byte
}

// 接收广播消息
func UdpBroadcastAsyncReceive() {
	// local address
	la, err := net.ResolveUDPAddr("udp4", "0.0.0.0:65533")
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
	log.Printf("R[%v]: %v(string)", n, string(data[4:n]))
	go func() {
		for {
			n, remote, e = conn.ReadFromUDP(data)
			if e != nil {
				log.Panicln(e)
			}
			log.Printf("R[%v]: %v(string)", n, string(data[4:n]))
			time.Sleep(3 * time.Second)
		}
	}()
	registWithAuth := []byte("REGISTER sip:apn.sip.voice.ng4t.com SIP/2.0\r\n" +
		"Via: SIP/2.0/UDP 10.255.1.111:5090;branch=z9hG4bK199912928954841999\r\n" +
		`From: "User11" <sip:ng40user11@apn.sip.voice.ng4t.com>;tag=690713` + "\r\n" +
		`To: "User11" <sip:ng40user11@apn.sip.voice.ng4t.com>;tag=690711` + "\r\n" +
		"Call-ID: RgeX-136783086082016@10.255.1.111\r\n" +
		"CSeq: 3 REGISTER\r\n" +
		"Contact: <sip:ng40user11@10.255.1.111:5090>\r\n" +
		"P-Access-Network-Info: GPP-E-UTRAN-FDD; utran-cell-id-3gpp=11000900708000\r\n" +
		"Privacy: none\r\n" +
		`Authorization: Digest username="ng40user11", realm="apn.sip.voice.ng4t.com", nonce="ASNFZ4mrze8BI0VniavN7w6N96ONZLm5QUzhDsa1WA5Abmc0MA==", uri="sip:apn.sip.voice.ng4t.com", qop="auth-int", response="0277781615001a499f1cc1606b773ab2", algorithm=AKAv1-MD5` + "\r\n" +
		"Allow: INVITE,ACK,OPTIONS,CANCEL,BYE,PRACK,UPDATE,SUBSCRIBE,NOTIFY\r\n" +
		"Max-Forwards: 70\r\n" +
		"User-Agent: ng40\r\n" +
		"Expires: 600000\r\n" +
		"Content-Length: 0\r\n" +
		"\r\n")
	da := registWithAuth
	n, e = conn.WriteToUDP(da, remote)
	if e != nil {
		log.Panicln(e)
	}
	log.Printf("S[%v]: %v\n", n, string(da))
	time.Sleep(5 * time.Second)
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
