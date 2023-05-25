package NKIOLC
/*
#cgo LDFLAGS: -L../../extlib -lNKIOLCLIB -lpthread -lstdc++
#cgo CFLAGS: -I../../extlib
#include "NK_IO_LC_LIB.h"
*/
import "C"
import "fmt"

import "it.etg/gpioServer/dto/NKIOLCDto"
import "it.etg/gpioServer/dto"

type GpioNKIOLC struct {
	Initialized bool
}

func NewGpioNKIOLC() *GpioNKIOLC{
	gpio := new(GpioNKIOLC)
	gpio.Initialized = false

	return gpio
}

func (gpio *GpioNKIOLC) Init() bool {
	path := "./config.ini"
    cpath := C.CString(path)
	res := C.DIOLC_LibraryBaseInit(cpath)

	if res <= 0 {
		panic("Errore inizializzazione gpio")
	} else {
		gpio.Initialized = true
	}

	return gpio.Initialized
}

func (gpio *GpioNKIOLC) Read() (bool, dto.GpioDto) {

	if !gpio.Initialized {
		return false, nil
	}

	diByte0 := byte(C.DIO_PollingReadDiByte(0));
	diByte1 := byte(C.DIO_PollingReadDiByte(1));

	fmt.Println(diByte0)
	fmt.Println(diByte1)

	var data NKIOLCDto.GpioNKIOLCDtoStruct

	for i:=0; i<8; i++ {
		if (diByte0 & (1 << i)) == 0 {
			data.X[i] = 0
		} else {
			data.X[i] = 1
		}
	}
	for i:=0; i<8; i++ {
		if (diByte1 & (1 << i)) == 0 {
			data.X[8+i] = 0
		} else {
			data.X[8+i] = 1
		}
	}

	return true, data
}

func (gpio *GpioNKIOLC) Write(gpioDto dto.GpioDto) bool {
	fmt.Println("Scrittura Digitali")
	fmt.Println(gpioDto)

	return true
}