package driver

import "it.etg/gpioServer/dto"

type Gpio interface {
	Init() bool
	Read() (bool, dto.GpioDto)
	Write(dto.GpioDto, dto.GpioWriteDto) bool
}