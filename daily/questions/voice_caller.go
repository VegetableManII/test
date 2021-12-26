package questions

/*
  #include <stdio.h>
  #include <unistd.h>
  #include <termios.h>
  char getch(){
      char ch = 0;
      struct termios old = {0};
      fflush(stdout);
      if( tcgetattr(0, &old) < 0 ) perror("tcsetattr()");
      old.c_lflag &= ~ICANON;
      old.c_lflag &= ~ECHO;
      old.c_cc[VMIN] = 1;
      old.c_cc[VTIME] = 0;
      if( tcsetattr(0, TCSANOW, &old) < 0 ) perror("tcsetattr ICANON");
      if( read(0, &ch,1) < 0 ) perror("read()");
      old.c_lflag |= ICANON;
      old.c_lflag |= ECHO;
      if(tcsetattr(0, TCSADRAIN, &old) < 0) perror("tcsetattr ~ICANON");
      return ch;
  }
*/
import "C"

import (
	"fmt"
	"net"
	"os"

	"github.com/gordonklaus/portaudio"
)

func VoiceCaller() {
	ra, err := net.ResolveUDPAddr("udp4", "127.0.0.1:7775")
	chk(err)
	conn, err := net.DialUDP("udp4", nil, ra)
	chk(err)
	chk(portaudio.Initialize())

	framesPerBuffer := make([]byte, 64*64)

	stream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(framesPerBuffer), framesPerBuffer)
	chk(err)

	go ExitWithESC(stream)

	// start reading from microphone
	errCheck(stream.Start())
	for {
		errCheck(stream.Read())
		fmt.Printf("\rListening...")
		// write to wave file
		_, err := conn.Write([]byte(framesPerBuffer))
		chk(err)
	}
}

func ExitWithESC(stream *portaudio.Stream) {
	key := C.getch()
	fmt.Println()
	fmt.Println("Cleaning up ...")
	if key == 27 {
		// better to control
		// how we close then relying on defer
		stream.Stop()
		stream.Close()
		portaudio.Terminate()
		os.Exit(0)

	}
}
