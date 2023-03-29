package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

const (
	MI          = 2.40466
	SAMPLE_RATE = 41000
)

type ToneGenerator struct {
	deviation float64
	frequency float64
	// serialPort  string
	// serialSpeed int
	sampleRate beep.SampleRate
	// pttType     int
}

func CreateToneGenerator() *ToneGenerator {
	tg := &ToneGenerator{}
	tg.SetDeviation("2200")
	tg.sampleRate = beep.SampleRate(SAMPLE_RATE)

	sr := beep.SampleRate(tg.sampleRate)
	speaker.Init(sr, sr.N(time.Second/10))

	return tg
}

func (tg *ToneGenerator) SetDeviation(s string) error {
	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}

	tg.deviation = value
	tg.frequency = value / MI

	return nil
}
func (tg *ToneGenerator) GetDevation() string {
	return fmt.Sprintf("%.0f", tg.deviation)
}

func (tg *ToneGenerator) GetFrequency() string {
	return fmt.Sprintf("%.0f", tg.frequency)
}

func (tg *ToneGenerator) Start() {
	sine, _ := CreateTone(tg.sampleRate, tg.frequency)

	speaker.Play(sine)
}

func (tg *ToneGenerator) Stop() {
	speaker.Clear()
}

func (tg *ToneGenerator) PTT() {
	// port, err := SerialInit(tg.serialPort, tg.serialSpeed)
	// if err != nil {
	// 	os.Exit(1)
	// }
	// defer port.Close()

	// err = port.SetRTS(true)
	// if err != nil {
	// 	fmt.Println("Error setting RTS:", err)
	// 	return
	// }

	// SetTX(port)

	// ClearTX(port)
}
