package main

import (
	"fmt"
	"math"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	// "github.com/faiface/beep/wav"
)

const (
	duration  = 1 * time.Second
	amplitude = 1.0 // Max amplitude is 1.0
)

func createTone(format beep.Format, frequency float64) (buffer *beep.Buffer) {
	// Create a buffer for the audio data
	buffer = beep.NewBuffer(format)

	// Calculate the number of samples needed for the tone duration
	numSamples := int(format.SampleRate.N(duration))

	sampNo := 0
	tone := beep.StreamerFunc(func(samples [][2]float64) (n int, ok bool) {
		if sampNo >= numSamples {
			return 0, false
		}

		fmt.Printf("StreamerFunc: %d\n", sampNo)
		for i := range samples {
			phase := float64(sampNo) / float64(format.SampleRate) * frequency * 2.0 * math.Pi
			samples[i][0] = amplitude * math.MaxInt16 * math.Sin(phase)
			samples[i][1] = samples[i][0]
			sampNo++
			if sampNo >= numSamples {
				break
			}
		}
		return len(samples), true
	})

	buffer.Append(tone)

	return
}

func main() {
	frequency := 1000.0 // Hz

	format := beep.Format{
		SampleRate:  beep.SampleRate(44100),
		NumChannels: 1,
		Precision:   2,
	}

	buffer := createTone(format, frequency)

	// Create a speaker and stream the audio data to it
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(buffer.Streamer(0, buffer.Len()))
	defer speaker.Close()

	select {}
}
