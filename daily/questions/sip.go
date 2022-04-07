package questions

import (
	"log"
	"strings"
	"study/daily/questions/sip"
)

func ViaHeader() {
	CRLF := "\r\n"
	sip.ServerDomain = "haha"
	msg, _ := sip.NewMessage(strings.NewReader("REGISTER sip:jiqimao@hebeiyidong.3gpp.net SIP/2.0" + CRLF +
		"Via: SIP/2.0/UDP 10.255.1.111:5090;branch=z9hG4bK199912928954841999" + CRLF + // 注册请求携带一个自己的VIP
		`From: "jiqimao" <sip:jiqimao@hebeiyidong.3gpp.net>;tag=690713` + CRLF +
		`To: "jiqimao" <sip:jiqimao@hebeiyidong.3gpp.net>;tag=690711` + CRLF + // 注册请求填自己
		"Call-ID: RgeX-136783086082016@10.255.1.111" + CRLF + // 随便编一个 目前网络侧没有用到
		"CSeq: 3 INVITE" + CRLF + // 客户端保证序列
		// "Contact: <sip:jiqimao@10.255.1.111:5090>" + CRLF + INVITE时需要填写自己实际局域网IP和端口
		"P-Access-Network-Info: 100231511300031" + CRLF +
		"Max-Forwards: 70" + CRLF +
		"Expires: 600000" + CRLF +
		"Content-Length: 0" + CRLF + CRLF))
	log.Println(msg.String())
	msg.Header.Via.AddServerInfo()
	log.Println("add", msg.String())
	msg.Header.Via.RemoveFirst()
	log.Println("rm", msg.String())
	log.Println(msg.String())
}
