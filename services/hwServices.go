package services

import "it.etg/gpioServer/libs/NKIOLCLibs"
import "it.etg/gpioServer/context"

func HwInit () {
	hwLibs := NKIOLCLibs.NewHwLibs()

	cx := context.GetContext()
	cx.HwLibs = hwLibs
}