package questions

import (
	"bufio"
	"log"
	"net"
	"os"
)

func Client_Main() {
	conn, err := net.Dial("tcp", "127.0.0.1:4399")
	if err != nil {
		log.Fatalln("[DEBUG]dial error ", err)
	}
	log.Println("[DEBUG]connected :4399")
	// net.DialTCP() 指定网络只能是tcp，指定本地地址为空随机选择，指定服务端主机地址
	stdin := bufio.NewReader(os.Stdin)
	cmd, err := stdin.ReadString('\n')
	if err != nil {
		log.Fatalln("[DEBUG]stdin read error ", err)
	}
	log.Println("[DEBUG]user input ", cmd)
	conn.Write([]byte(cmd))

	log.Println("[DEBUG]write into net ", cmd)
	stdout := bufio.NewWriter(os.Stdout)
	connScanner := bufio.NewScanner(conn)
	for {
		if connScanner.Scan() {
			str := connScanner.Text()
			log.Println(str)
			stdout.Write([]byte(str))
		}
	}
}
