package questions

import (
	"log"
	"net"
)

func UdpBroadcastServer() {
	l, e := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4bcast,
		Port: 20379,
	})
	if e != nil {
		log.Panicln(e)
	}
	log.Println("server listen on", l.LocalAddr().String(), l.LocalAddr().Network())
	data := make([]byte, 0, 512)
	for {
		n, remote, e := l.ReadFromUDP(data)
		if e != nil {
			log.Panicln(e)
		}
		log.Println("client", remote, n)
		log.Println("msg", data)
	}
}
