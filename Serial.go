package main

import (
	"log"

	"go.bug.st/serial"
)

func SerialInit(serialPort string, speed int) (serial.Port, error) {
	if serialPort == "" {
		return nil, nil
	}

	mode := &serial.Mode{
		BaudRate:          speed,
		DataBits:          8,
		Parity:            serial.NoParity,
		InitialStatusBits: &serial.ModemOutputBits{RTS: false, DTR: false},
	}

	port, err := serial.Open(serialPort, mode)
	if err != nil {
		log.Fatal(err)
	}

	return port, nil
}

func SetTX(port serial.Port) {
	if port == nil {
		return
	}

	err := port.SetRTS(true)
	if err != nil {
		log.Fatal("SetRTS(true) failed: ", err)
	}
	err = port.SetDTR(true)
	if err != nil {
		log.Fatal("SetDTR(true) failed: ", err)
	}
}

func ClearTX(port serial.Port) {
	if port == nil {
		return
	}

	err := port.SetRTS(false)
	if err != nil {
		log.Fatal("SetRTS(true) failed: ", err)
	}
	err = port.SetDTR(false)
	if err != nil {
		log.Fatal("SetDTR(true) failed: ", err)
	}
}
