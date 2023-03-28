package main

import (
	"fmt"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

const MI = 2.40466

func main() {
	sr := beep.SampleRate(44100)

	deviation := 2200.0
	freq := deviation / MI
	fmt.Printf("Freq: %f", freq)
	sine, _ := CreateTone(sr, freq)

	speaker.Init(sr, sr.N(time.Second/10)) // sr.N(time.Second/10) = buffer size for duration 1/10 second
	speaker.Play(sine)
	select {} // makes the program hang forever
}
