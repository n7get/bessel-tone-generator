package main

import (
	"go.bug.st/serial"
)

var port serial.Port

func SerialOpen(path string, speed int) error {
	mode := &serial.Mode{
		BaudRate:          speed,
		DataBits:          8,
		Parity:            serial.NoParity,
		InitialStatusBits: &serial.ModemOutputBits{RTS: false, DTR: false},
	}

	var err error
	port, err = serial.Open(path, mode)
	if err != nil {
		return err
	}

	return nil
}

func SetRTS(value bool) error {
	if err := port.SetRTS(value); err != nil {
		return err
	}
	return nil
}

func SetDTR(value bool) error {
	if err := port.SetDTR(value); err != nil {
		return err
	}
	return nil
}

func SerialClose() {
	port.Close()
	port = nil
}

func SerialPortList() ([]string, error) {
	return serial.GetPortsList()
}
