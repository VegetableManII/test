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

// 接收广播消息
func UdpBroadcastAsyncReceive() {
	quit := make(chan os.Signal, 1)
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

	go func() {
		for {
			data := make([]byte, 1024)
			n, remote, e = conn.ReadFromUDP(data)
			if e != nil {
				log.Panicln(e)
			}
			log.Printf("R[%v]: %v(string)", n, string(data[:n]))
			time.Sleep(3 * time.Second)
		}
	}()
	CRLF := "\r\n"
	registWithAuth := []byte("INVITE sip:jiqimao@hebeiyidong.3gpp.net SIP/2.0" + CRLF +
		"Via: SIP/2.0/UDP 10.255.1.111:5090;branch=z9hG4bK199912928954841999" + CRLF + // 注册请求携带一个自己的VIP
		`From: "jiqimao" <sip:jiqimao@hebeiyidong.3gpp.net>;tag=690713` + CRLF +
		`To: "jiqimao" <sip:jiqimao@hebeiyidong.3gpp.net>;tag=690711` + CRLF + // 注册请求填自己
		"Call-ID: RgeX-136783086082016@10.255.1.111" + CRLF + // 随便编一个 目前网络侧没有用到
		"CSeq: 3 INVITE" + CRLF + // 客户端保证序列
		// "Contact: <sip:jiqimao@10.255.1.111:5090>" + CRLF + INVITE时需要填写自己实际局域网IP和端口
		"P-Access-Network-Info: 100231511300031" + CRLF +
		"Max-Forwards: 70" + CRLF +
		"Expires: 600000" + CRLF +
		"Content-Length: 0" + CRLF + CRLF)

	<-quit
	msg := []byte(`{
		"protocal":"epc",
		"method":"attach request",
		"utan-cell-id-3gpp": "100231511300031"
	}`)
	n, e = conn.WriteToUDP(msg, remote)
	_ = registWithAuth
	// n, e = conn.WriteToUDP(registWithAuth, remote)
	if e != nil {
		log.Panicln(e)
	}

	<-quit
}

func stringToBytes(s string) []byte {
	b := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		b = append(b, s[i]-0x30)
	}
	return b
}

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

/*     协议(1B)方法(1B)长度(2B)
	字节  0     1     2     3
        0x01, 0x00, 0x00, 0x22
		UTRAN-CELL-ID-3GPP=100231511300031

		协议(1B)方法(1B)长度(2B)
	字节  0     1     2     3
        0x01, 0x0A, 0x00, 0x22  (下面消息长度的的16进制标识)
		UTRAN-CELL-ID-3GPP=100231511300031
		IP=10.10.10.1
*/
