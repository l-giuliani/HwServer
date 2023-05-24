package context

import "it.etg/gpioServer/driver"

type Context struct {
	GpioDrv driver.Gpio
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