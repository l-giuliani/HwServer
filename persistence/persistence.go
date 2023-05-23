package persistence

import "it.etg/gpioServer/dto"

type Persistence struct {
	GpioData dto.GpioDto
}

var persistence Persistence

func GetPersistence() *Persistence{
	return &persistence
}