package services

import "it.etg/gpioServer/driver/NKIOLC"
import "it.etg/gpioServer/dao"
import "it.etg/gpioServer/dto"
import "it.etg/gpioServer/context"
import "fmt"

func GpioInit() {
	gpioDrv := NKIOLC.NewGpioNKIOLC()
	//gpioDrv.Init()

	cx := context.GetContext()
	cx.GpioDrv = gpioDrv

	GpioRead()
}

func GpioContinuousRead() {
	
}

func StopContinuousRead() {

}

func GpioRead() {
	cx := context.GetContext()
	res, gpioDto := cx.GpioDrv.Read()
	if !res {
		fmt.Println("Errore Lettura Gpio")
		return
	}
	fmt.Println(gpioDto)
	dao.SetGpioData(gpioDto)
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
	return cx.HwLibs.ParseData(data)
}

func GpioWrite(data []byte) bool {
	res, gpioDto := DecodeGpioData(data)
	if !res {
		return false
	}
	cx := context.GetContext()
	return cx.GpioDrv.Write(gpioDto)
}