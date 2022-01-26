package questions

import (
	"log"
	"net"
)

func GetNetInterfaces() {
	ifs, err := net.Interfaces()
	if err != nil {
		log.Panicln(err)
	}
	for _, v := range ifs {
		log.Println(v)
		log.Println(v.Flags.String(), v.HardwareAddr.String())
		addrs, _ := v.Addrs()
		log.Println(len(addrs))
		for _, v := range addrs {
			log.Println(v.Network(), v.String())
		}
		multiaddrs, _ := v.Addrs()
		log.Println(len(multiaddrs))
		for _, v := range multiaddrs {
			log.Println(v.Network(), v.String())
		}
	}
}
