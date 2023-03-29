package main

import (
	"fmt"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

const MI = 2.40466

func main() {
	deviation := 2200.0
	var serialPort string
	speed := 38400
	sampleRate := 44100

	sr := beep.SampleRate(sampleRate)
	freq := deviation / MI
	fmt.Printf("Freq: %f", freq)
	sine, _ := CreateTone(sr, freq)

	port, err := SerialInit(serialPort, speed)
	if err != nil {
		os.Exit(1)
	}
	defer port.Close()

	err = port.SetRTS(true)
	if err != nil {
		fmt.Println("Error setting RTS:", err)
		return
	}

	SetTX(port)

	speaker.Init(sr, sr.N(time.Second/10)) // sr.N(time.Second/10) = buffer size for duration 1/10 second
	speaker.Play(sine)
	select {} // makes the program hang forever

	// ClearTX(port)
}
