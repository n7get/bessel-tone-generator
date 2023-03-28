package main

import "math"

type SineWave struct {
	sampleFactor float64 // Just for ease of use so that we don't have to calculate every sample
	phase        float64
}

func (g *SineWave) Stream(samples [][2]float64) (n int, ok bool) {
	for i := range samples { // increment = ((2 * PI) / SampleRate) * freq
		v := math.Sin(g.phase * 2.0 * math.Pi) // period of the wave is thus defined as: 2 * PI.
		samples[i][0] = v
		samples[i][1] = v
		_, g.phase = math.Modf(g.phase + g.sampleFactor)
	}

	return len(samples), true
}

func (*SineWave) Err() error {
	return nil
}
