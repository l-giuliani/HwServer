package libs

import "it.etg/gpioServer/dto"

type HwLibs interface {
	ParseWriteData([]byte) (bool, dto.GpioWriteDto)
	GetFormattedData(dto.GpioDto) []byte
	CheckGpioDataChanged(dto.GpioDto, dto.GpioDto) bool
	GetOutputDataByDo(output string) dto.GpioWriteDto
	GetResetDataByWriteDto(dto.GpioWriteDto) []dto.ResetTimeData
}