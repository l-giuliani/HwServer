package services

import "it.etg/gpioServer/driver/NKIOLC"
import "it.etg/gpioServer/context"

func StartWd(time uint8) bool{
	cx := context.GetContext()
	return cx.WdDrv.StartWd(time)
}

func WdInit() {
	wdDrv := NKIOLC.NewWdNKIOLC()
	wdDrv.Init()

	cx := context.GetContext()
	cx.WdDrv = wdDrv

	if cx.Conf.WdActive {
		StartWd(cx.Conf.WdStartTime)
	}
}

func ResetWd(time uint8) bool {
	cx := context.GetContext()
	return cx.WdDrv.ResetWd(time)
}

func StopWd() bool {
	cx := context.GetContext()
	return cx.WdDrv.StopWd()
}