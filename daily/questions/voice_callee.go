package questions

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/gordonklaus/portaudio"
)

func VoiceCallee() {
	la, err := net.ResolveUDPAddr("udp4", "127.0.0.1:7775")
	chk(err)
	conn, err := net.ListenUDP("udp4", la)
	chk(err)
	chk(portaudio.Initialize())
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	framesPerBuffer := make([]byte, 64*64)
	stream, err := portaudio.OpenDefaultStream(0, 1, 44100, len(framesPerBuffer), framesPerBuffer)
	chk(err)

	go ExitWithESC(stream)
	chk(stream.Start())
	log.Println("stream.Start")
	for {
		_, _, err := conn.ReadFromUDP(framesPerBuffer)
		chk(err)
		for i := range framesPerBuffer {
			framesPerBuffer[i] = byte(0.7 * float32(framesPerBuffer[i]))
		}
		stream.Write()
	}
}
