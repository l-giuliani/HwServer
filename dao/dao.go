package dao

import "it.etg/gpioServer/dto"
import "it.etg/gpioServer/persistence"

func GetGpioData() dto.GpioDto {
	persistence := persistence.GetPersistence()

	return persistence.GpioData
} 

func SetGpioData(gpioData dto.GpioDto) {
	persistence := persistence.GetPersistence()
	persistence.GpioData = gpioData
}