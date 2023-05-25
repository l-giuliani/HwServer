package NKIOLC
/*
#cgo LDFLAGS: -L../../extlib -lNKIOLCLIB -lpthread -lstdc++
#cgo CFLAGS: -I../../extlib
#include "NK_IO_LC_LIB.h"
*/
import "C"
import "fmt"
import "reflect"

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

	doByte0 := byte(C.DIO_PollingReadDoByte(0));
	doByte1 := byte(C.DIO_PollingReadDoByte(1));

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

	for i:=0; i<8; i++ {
		if (doByte0 & (1 << i)) == 0 {
			data.Y[i] = 0
		} else {
			data.Y[i] = 1
		}
	}
	for i:=0; i<8; i++ {
		if (doByte1 & (1 << i)) == 0 {
			data.Y[8+i] = 0
		} else {
			data.Y[8+i] = 1
		}
	}

	return true, data
}

func (gpio *GpioNKIOLC) Write(gpioDto dto.GpioDto, gpioWriteDto dto.GpioWriteDto) bool {
	fmt.Println("Scrittura Digitali")
	fmt.Println(gpioWriteDto)
	gpioSpec := gpioDto.(NKIOLCDto.GpioNKIOLCDtoStruct)
	gpioWriteSpec := gpioWriteDto.(NKIOLCDto.GpioNKIOLCSetData)

	var newByte0 byte
	newByte0 = 0
	var newByte1 byte
	newByte1 = 0

	for i:=0; i<8; i++  {
		newByte0 = newByte0 | (gpioSpec.Y[i] << i)
	}
	for i:=0; i<8; i++  {
		newByte1 = newByte1 | (gpioSpec.Y[8+i] << i)
	}

	v := reflect.ValueOf(gpioWriteSpec)

	for i := 0; i< v.NumField(); i++ {
        el := v.Field(i).Interface().(NKIOLCDto.GpioNKIOLCSetDataEl)
		if !el.Update {
			continue
		}
		if i < 8 {
			if el.Value == 0 {
				newByte0 = newByte0 &^ (1 << i)
			} else {
				newByte0 = newByte0 | (1 << i)
			}
			
		} else {
			if el.Value == 0 {
				newByte1 = newByte1 &^ (1 << (i-8))
			} else {
				newByte1 = newByte1 | (1 << (i-8))
			}
		}
    }

	C.DIO_PollingWriteDoByte(0, C.uchar(newByte0));
	C.DIO_PollingWriteDoByte(1, C.uchar(newByte1));

	return true
}