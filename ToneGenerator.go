package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/speaker"
	"golang.org/x/exp/slices"
)

const (
	MI          = 2.40466
	SAMPLE_RATE = 41000
)

var SERIAL_SPEEDS = []int64{4800, 9600, 19200, 38400, 56700, 115200}

const (
	PTT_NONE = iota
	PTT_RTS
	PTT_DTR
	PTT_BOTH
)

type ToneGenerator struct {
	deviation   float64
	frequency   float64
	level       float64
	serialPort  string
	serialSpeed int
	sampleRate  beep.SampleRate
	pttType     int
}

func CreateToneGenerator() *ToneGenerator {
	tg := &ToneGenerator{
		level:       0.67,
		serialSpeed: 9600,
		pttType:     PTT_NONE,
	}
	tg.SetDeviation("2200")
	tg.sampleRate = beep.SampleRate(SAMPLE_RATE)

	sr := beep.SampleRate(tg.sampleRate)
	speaker.Init(sr, sr.N(time.Second/10))

	return tg
}

func (tg *ToneGenerator) GetDevation() string {
	return fmt.Sprintf("%.0f", tg.deviation)
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

func (tg *ToneGenerator) GetFrequency() string {
	return fmt.Sprintf("%.0f", tg.frequency)
}

func (tg *ToneGenerator) GetLevel() string {
	return fmt.Sprintf("%.02f", tg.level)
}
func (tg *ToneGenerator) SetLevel(s string) error {
	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}

	tg.level = value

	return nil
}

func (tg *ToneGenerator) GetPttType() string {
	switch tg.pttType {
	case PTT_NONE:
		return "NONE"
	case PTT_RTS:
		return "RTS"
	case PTT_DTR:
		return "DTR"
	case PTT_BOTH:
		return "BOTH"
	}
	return "UNKNOWN"
}
func (tg *ToneGenerator) SetPttType(value string) {
	switch value {
	case "NONE":
		tg.pttType = PTT_NONE
	case "RTS":
		tg.pttType = PTT_RTS
	case "DTR":
		tg.pttType = PTT_DTR
	case "BOTH":
		tg.pttType = PTT_BOTH
	}
}

func (tg *ToneGenerator) GetSerialPort() string {
	return tg.serialPort
}
func (tg *ToneGenerator) SetSerialPort(value string) {
	tg.serialPort = value
}

func (tg *ToneGenerator) GetSerialSpeed() string {
	return fmt.Sprintf("%d", tg.serialSpeed)
}
func (tg *ToneGenerator) SetSerialSpeed(value string) error {
	speed, err := strconv.ParseInt(value, 10, 16)
	if err != nil {
		return err
	}

	if slices.Contains(SERIAL_SPEEDS, speed) {
		return errors.New("unknown serial port speed")
	}

	tg.serialSpeed = int(speed)

	return nil
}

func (tg *ToneGenerator) Start() error {
	tg.PttOn()

	sine, _ := CreateTone(tg.sampleRate, tg.frequency)
	level := &effects.Gain{
		Streamer: sine,
		Gain:     -(1 - tg.level),
	}
	speaker.Play(level)

	return nil
}

func (tg *ToneGenerator) Stop() {
	tg.PttOff()
	speaker.Clear()
}

func (tg *ToneGenerator) PttOn() error {
	if tg.pttType == PTT_NONE {
		return nil
	}

	if err := SerialOpen(tg.serialPort, tg.serialSpeed); err != nil {
		return err
	}

	if tg.pttType == PTT_RTS || tg.pttType == PTT_BOTH {
		if err := SetRTS(true); err != nil {
			SerialClose()
			return err
		}
	}
	if tg.pttType == PTT_DTR || tg.pttType == PTT_BOTH {
		if err := SetDTR(true); err != nil {
			SerialClose()
			return err
		}
	}

	return nil
}
func (tg *ToneGenerator) PttOff() error {
	if tg.pttType == PTT_RTS || tg.pttType == PTT_BOTH {
		SetRTS(false)
	}
	if tg.pttType == PTT_DTR || tg.pttType == PTT_BOTH {
		SetDTR(false)
	}

	if tg.pttType != PTT_NONE {
		SerialClose()
	}

	return nil
}

func (tg *ToneGenerator) CanStart() bool {
	return tg.pttType == PTT_NONE || len(tg.serialPort) > 0
}
