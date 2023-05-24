package context

import "it.etg/gpioServer/driver"
import "it.etg/gpioServer/libs"

type Context struct {
	GpioDrv driver.Gpio
	HwLibs	libs.HwLibs
}

var context Context
var initialized bool = false

func Init() {
	context = Context{}
	initialized = true
}

func GetContext() *Context {
	if initialized {
		return &context
	} else {
		return nil
	}
}