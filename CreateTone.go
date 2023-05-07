package main

import (
	"errors"

	"github.com/faiface/beep"
)

func CreateTone(sr beep.SampleRate, freq float64, level float64) (beep.Streamer, error) {
	dt := freq / float64(sr)

	if dt >= 1.0/2.0 {
		return nil, errors.New("samplerate must be at least 2 times grater then frequency")
	}

	return &SineWave{dt, 0.1, level}, nil
}
