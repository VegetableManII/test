package questions

// stackoverflow.com/questions/14094190/golang-function-similar-to-getchar

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/gordonklaus/portaudio"
	wave "github.com/zenwerk/go-wave"
)

func errCheck(err error) {

	if err != nil {
		panic(err)
	}
}

func RecordMicrophoneToWAV() {

	if len(os.Args) != 2 {
		fmt.Printf("Usage : %s <audiofilename.wav>\n", os.Args[0])
		os.Exit(0)
	}

	audioFileName := os.Args[1]

	fmt.Println("Recording. Press ESC to quit.")

	if !strings.HasSuffix(audioFileName, ".wav") {
		audioFileName += ".wav"
	}
	waveFile, err := os.Create(audioFileName)
	errCheck(err)

	// www.people.csail.mit.edu/hubert/pyaudio/  - under the Record tab
	inputChannels := 1
	outputChannels := 0
	sampleRate := 44100
	framesPerBuffer := make([]byte, 64*64)

	// init PortAudio

	portaudio.Initialize()
	//defer portaudio.Terminate()

	stream, err := portaudio.OpenDefaultStream(inputChannels, outputChannels, float64(sampleRate), len(framesPerBuffer), framesPerBuffer)
	errCheck(err)
	//defer stream.Close()

	// setup Wave file writer

	param := wave.WriterParam{
		Out:           waveFile,
		Channel:       inputChannels,
		SampleRate:    sampleRate,
		BitsPerSample: 8, // if 16, change to WriteSample16()
	}

	waveWriter, err := wave.NewWriter(param)
	errCheck(err)

	//defer waveWriter.Close()

	go ExitWithESC(stream)

	// recording in progress ticker. From good old DOS days.
	ticker := []string{
		"-",
		"\\",
		"/",
		"|",
	}
	rand.Seed(time.Now().UnixNano())

	// start reading from microphone
	errCheck(stream.Start())
	for {
		errCheck(stream.Read())

		fmt.Printf("\rRecording is live now. Say something to your microphone! [%v]", ticker[rand.Intn(len(ticker)-1)])

		// write to wave file
		_, err := waveWriter.Write([]byte(framesPerBuffer)) // WriteSample16 for 16 bits
		errCheck(err)
	}
	errCheck(stream.Stop())
}
