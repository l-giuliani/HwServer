package services

import "it.etg/gpioServer/driver/NKIOLC"
import "it.etg/gpioServer/dao"
import "fmt"

func GpioContinuousRead() {
	
}

func StopContinuousRead() {

}

func GpioRead() {
	gpioDrv := NKIOLC.NewGpioNKIOLC()
	gpioDrv.Init()
	res, gpioDto := gpioDrv.Read()
	if !res {
		fmt.Println("Errore Lettura Gpio")
		return
	}
	fmt.Println(gpioDto)
	dao.SetGpioData(gpioDto)
	
}