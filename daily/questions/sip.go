package questions

import (
	"strings"
	"study/daily/questions/sip"
)

func ViaHeader() {
	sip.ServerDomain = "haha"
	msg, _ := sip.NewMessage(strings.NewReader(`REGISTER sip:192.168.0.2:5060 SIP/2.0
	Via: SIP/2.0/UDP 192.168.0.102:5060;branch=z9hG4bK-172441109;rport
	From: "jiqimao" <sip:jiqimao@hebeiyidong.3gpp.net>;tag=690713
	To: "jiqimao" <sip:jiqimao@hebeiyidong.3gpp.net>;tag=690711
	Call-ID: 012022699881-214NTIL@192.168.0.102
	CSeq: 1 REGISTER
	Expires: 600
	Allow: INVITE,CANCEL,ACK,BYE,NOTIFY,REFER,OPTIONS,INFO,MESSAGE,UPDATE,PRACK
	Max-Forwards: 70
	User-Agent: NTIL214 SIP-Test UA
	Content-Length:0
	Contact: <sip:1010@192.168.0.102:5060>`))
	msg.Header.Via.AddServerInfo()
	msg.Header.Via.RemoveFirst()
}
