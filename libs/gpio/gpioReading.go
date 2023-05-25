package libs

import "time"
import "it.etg/gpioServer/context"
import "it.etg/gpioServer/dao"
import "it.etg/gpioServer/dto"
import "it.etg/gpioServer/libs/network"
import "fmt"

var readExecuting bool = false
var udpClient *network.UdpClient = nil

func gpioContinuousRead() {
	start := time.Now()
	for readExecuting {
		t := time.Now()
		if t.Sub(start) >= time.Duration(context.GetContext().Conf.AcquTime) * 1000000000 {
			res, gpioData := GpioRead()
			if res {
				NotifyStatus(gpioData)
				dao.SetGpioData(gpioData)
			}
			start = t
		}
		time.Sleep(time.Second)
	}
	
}

func NotifyStatus(gpioData dto.GpioDto) {
	if context.GetContext().Conf.UDPEnabled {
		if checkGpioStatusChanged(gpioData) {
			if udpClient == nil {
				udpClient = network.NewUdpClient()
			}
			udpClient.WriteAndClose(context.GetContext().Conf.UDPAddress, []byte("DO"))
		}
	}
}

func StartContinuousRead() {
	if !readExecuting {
		readExecuting = true
		go gpioContinuousRead()
	}
}

func StopContinuousRead() {
	readExecuting = false
}

func GpioRead() (bool, dto.GpioDto) {
	cx := context.GetContext()
	res, gpioDto := cx.GpioDrv.Read()
	if !res {
		fmt.Println("Errore Lettura Gpio")
		return false, gpioDto
	}
	fmt.Println(gpioDto)
	//dao.SetGpioData(gpioDto)

	return true, gpioDto
}

func checkGpioStatusChanged(data dto.GpioDto) bool {
	gpioDto := dao.GetGpioData()
	fmt.Println("=========")
	fmt.Println(data)
	fmt.Println(gpioDto)
	cx := context.GetContext()
	return cx.HwLibs.CheckGpioDataChanged(gpioDto, data)
}