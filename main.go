package main

import (
	"errors"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	"fyne.io/fyne/v2/widget"
)

func main() {
	tg := CreateToneGenerator()

	a := app.New()
	w := a.NewWindow("Bessel Tone Generator")

	startStop := widget.NewButton("Start", nil)
	startStop.OnTapped = func() {
		if startStop.Text == "Start" {
			startStop.SetText("Stop")
			tg.Start()
		} else {
			startStop.SetText("Start")
			tg.Stop()
		}
	}

	deviationEntry := widget.NewEntry()
	deviationEntry.SetText(tg.GetDevation())
	deviationEntry.Validator = func(text string) error {
		value, err := strconv.ParseFloat(text, 64)
		if err != nil {
			return errors.New("value not numeric")
		}
		if value < 2000 || value > 3000 {
			return errors.New("deviation value is out or range")
		}
		return nil
	}
	deviationEntry.OnChanged = func(s string) {
		fmt.Printf("OnChanged: %s\n", s)
		tg.SetDeviation(s)
	}

	serialPorts, _ := SerialPortList()
	serialDevEntry := widget.NewSelect(serialPorts, func(s string) {
		tg.SetSerialPort(s)
	})
	serialDevEntry.SetSelected(tg.GetSerialPort())

	serialSpeedEntry := widget.NewSelect([]string{"9600", "19200", "38400", "56700", "155200"}, func(s string) {
		tg.SetSerialSpeed(s)
	})
	serialSpeedEntry.SetSelected(tg.GetSerialSpeed())

	pttType := widget.NewSelect([]string{"NONE", "RTS", "DTR", "BOTH"}, func(s string) {
		tg.SetPttType(s)
	})
	pttType.SetSelected(tg.GetPttType())

	w.SetContent(container.NewVBox(
		container.New(
			layout.NewFormLayout(),
			widget.NewLabel("Deviation"),
			deviationEntry,
			widget.NewLabel("Serial Device"),
			serialDevEntry,
			widget.NewLabel("Serial Speed"),
			serialSpeedEntry,
			widget.NewLabel("PTT Type"),
			pttType,
		),
		layout.NewSpacer(),
		container.NewCenter(startStop),
	))

	w.ShowAndRun()
}
