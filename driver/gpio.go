package driver

import "it.etg/gpioServer/dto"

type Gpio interface {
	Init() bool
	Read() (uint16, dto.GpioDto) 
}