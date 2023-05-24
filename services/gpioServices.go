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
	fmt.Println(gpioData)

	return gpioData
}