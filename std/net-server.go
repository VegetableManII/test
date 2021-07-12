package main

import (
	"net"
)

func handleconnection(c net.Conn) {
	c.Write([]byte("hello"))
}

// func main() {
// 	srv, err := net.Listen("tcp", ":4411")
// 	if err != nil {
// 		log.Fatal("Create Server Error!")
// 	}
// 	for {
// 		conn, err := srv.Accept()
// 		if err != nil {
// 			log.Fatal("Connect Error")
// 		}
// 		go handleconnection(conn)
// 	}

// }

// func main() {
// 	// Listen on TCP port 2000 on all available unicast and
// 	// anycast IP addresses of the local system.
// 	l, err := net.Listen("tcp", ":2000")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer l.Close()
// 	for {
// 		// Wait for a connection.
// 		conn, err := l.Accept()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		// Handle the connection in a new goroutine.
// 		// The loop then returns to accepting, so that
// 		// multiple connections may be served concurrently.
// 		go func(c net.Conn) {
// 			// Echo all incoming data.
// 			io.Copy(c, c)
// 			// Shut down the connection.
// 			c.Close()
// 		}(conn)
// 	}
// }
