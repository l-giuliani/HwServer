package services

import "it.etg/gpioServer/driver/NKIOLC"
import "it.etg/gpioServer/dao"
import "it.etg/gpioServer/dto"
import "it.etg/gpioServer/libs/gpio"
import "it.etg/gpioServer/context"

func GpioInit(continuousRead bool) {
	gpioDrv := NKIOLC.NewGpioNKIOLC()
	gpioDrv.Init()

	cx := context.GetContext()
	cx.GpioDrv = gpioDrv

	GpioRead(false)

	if continuousRead {
		libs.StartContinuousRead()
	}
}

func GpioRead(notify bool) {
	res, gpioData := libs.GpioRead()
	if res {
		if notify {
			libs.NotifyStatus(gpioData)
		}
		dao.SetGpioData(gpioData)
	}
}

func GpioGetData() dto.GpioDto{
	gpioData := dao.GetGpioData()

	return gpioData
}

func FormatGpioData(gpioDto dto.GpioDto) []byte{
	cx := context.GetContext()
	return cx.HwLibs.GetFormattedData(gpioDto)
}

func DecodeGpioData(data []byte) (bool, dto.GpioDto) {
	cx := context.GetContext()
	return cx.HwLibs.ParseWriteData(data)
}

func GpioWrite(data []byte) bool {
	res, gpioWriteDto := DecodeGpioData(data)
	if !res {
		return false
	}
	libs.GpioWrite(gpioWriteDto, false)

	//gpioData := dao.GetGpioData()
	//cx := context.GetContext()
	//cx.GpioDrv.Write(gpioData, gpioWriteDto)
	GpioRead(true)
	return true
}