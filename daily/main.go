package main

import "study/daily/questions"

func main() {
	questions.UdpBroadcastAsyncReceive()
}

// var uriRexp = regexp.MustCompile("^(SIP\\/[^\\/]+)\\/([^ ]+) ([^;]+)(.+)$")

// func main() {
// 	str := "SIP/2.0/U DP TCP 192. 168.0.100:43188;branch=z9hG4bK111643fe9a9f389667c5e7d8873;rport"
// 	sub := uriRexp.FindStringSubmatch(str)
// 	for _, v := range sub {
// 		log.Println(v)
// 	}
// }
