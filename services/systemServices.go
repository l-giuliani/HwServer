package services

import "it.etg/gpioServer/context"
import "it.etg/gpioServer/config"

func configInit() {
	conf := config.NewConfig()
	conf.Init()
	context.GetContext().Conf = conf
}

func SystemInit() {
	context.Init()
	configInit()
	HwInit()
	GpioInit(true)
	WdInit()
}