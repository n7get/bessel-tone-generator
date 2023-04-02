package main

import (
	"errors"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const (
	MIN_DEVIATION = 1500
	MAX_DEVIATION = 4000
)

func CreateGui(tg *ToneGenerator) {

	a := app.New()
	w := a.NewWindow("Bessel Tone Generator")
	w.Resize(fyne.NewSize(300, 200))

	var startStop *widget.Button
	startStop = widget.NewButton("Start", func() {
		if startStop.Text == "Start" {
			startStop.SetText("Stop")
			startStop.Importance = widget.DangerImportance
			tg.Start()
		} else {
			startStop.SetText("Start")
			tg.Stop()
			startStop.Importance = widget.HighImportance
		}
	})
	startStop.Importance = widget.HighImportance

	deviationLabel := widget.NewLabel("Deviation")
	deviationEntry := widget.NewEntry()
	deviationEntry.SetText(tg.GetDevation())
	deviationEntry.Validator = func(text string) error {
		value, err := strconv.ParseFloat(text, 64)
		if err != nil {
			return errors.New("value not numeric")
		}
		if value < MIN_DEVIATION || value > MAX_DEVIATION {
			return errors.New("deviation value is out or range")
		}
		return nil
	}
	deviationEntry.OnChanged = func(s string) {
		tg.SetDeviation(s)
	}

	serialDevLabel := widget.NewLabel("Serial Device")
	serialPorts, _ := SerialPortList()
	serialDevEntry := widget.NewSelect(serialPorts, func(s string) {
		tg.SetSerialPort(s)
		if tg.CanStart() {
			startStop.Enable()
		} else {
			startStop.Disable()
		}
	})
	serialDevEntry.SetSelected(tg.GetSerialPort())

	// serialSpeedLabel := widget.NewLabel("Serial Speed")
	// serialSpeedEntry := widget.NewSelect([]string{"9600", "19200", "38400", "56700", "155200"}, func(s string) {
	// 	tg.SetSerialSpeed(s)
	// })
	// serialSpeedEntry.SetSelected(tg.GetSerialSpeed())

	pttTypeLabel := widget.NewLabel("PTT Type")
	pttTypeSelect := widget.NewSelect([]string{"NONE", "RTS", "DTR", "BOTH"}, func(s string) {
		tg.SetPttType(s)
		if s != "NONE" {
			serialDevEntry.Enable()
		} else {
			serialDevEntry.Disable()
		}
		if tg.CanStart() {
			startStop.Enable()
		} else {
			startStop.Disable()
		}
	})
	pttTypeSelect.SetSelected(tg.GetPttType())

	w.SetContent(container.NewVBox(
		container.New(
			layout.NewFormLayout(),
			deviationLabel,
			deviationEntry,
			pttTypeLabel,
			pttTypeSelect,
			serialDevLabel,
			serialDevEntry,
			// serialSpeedLabel,
			// serialSpeedEntry,
		),
		layout.NewSpacer(),
		container.NewCenter(startStop),
	))

	w.ShowAndRun()
}
