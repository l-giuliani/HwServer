package services

import "it.etg/gpioServer/context"

func SystemInit() {
	context.Init()
	GpioInit()
}