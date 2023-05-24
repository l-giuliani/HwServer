package libs

import "it.etg/gpioServer/dto"

type HwLibs interface {
	ParseData([]byte) (bool, dto.GpioDto)
	GetFormattedData(dto.GpioDto) []byte
}