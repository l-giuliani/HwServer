package libs

import "time"
import "it.etg/gpioServer/context"
import "it.etg/gpioServer/dao"
import "it.etg/gpioServer/dto"
import "it.etg/gpioServer/libs/network"
import "sync"
import "fmt"

var readExecuting bool = false
var udpClient *network.UdpClient = nil

var resetMutex sync.Mutex
var mutexLocked bool = false
var resetTimeData []dto.ResetTimeData
var initResetTimeData bool = false
var rwMutex sync.Mutex
var rsMutex sync.Mutex

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
			start = time.Now()
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

	rwMutex.Lock()
	res, gpioDto := cx.GpioDrv.Read()
	rwMutex.Unlock()

	if !res {
		fmt.Println("Errore Lettura Gpio")
		return false, gpioDto
	}
	//fmt.Println(gpioDto)
	//dao.SetGpioData(gpioDto)

	return true, gpioDto
}

func checkGpioStatusChanged(data dto.GpioDto) bool {
	gpioDto := dao.GetGpioData()
	cx := context.GetContext()
	return cx.HwLibs.CheckGpioDataChanged(gpioDto, data)
}

func ResetGpoStateFun() {
	cx := context.GetContext()
	for true {
		if len(resetTimeData) == 0 {
			mutexLocked = true
			resetMutex.Lock()
			continue
		}
		now := time.Now()
		changed := false

		rsMutex.Lock()
		for i:=0;i<len(resetTimeData);i++ {
			if now.Sub(resetTimeData[i].StartTime) >= time.Duration(resetTimeData[i].Time) * 1000000000 {
				gpoData := cx.HwLibs.GetOutputDataByDo(resetTimeData[i].Output)
				GpioWrite(gpoData, true)
				changed = true
				resetTimeData = append(resetTimeData[:i], resetTimeData[i+1:]...)
				i--
			}
		}
		rsMutex.Unlock()

		if changed {
			res, gpioR := GpioRead()
			if res {
				NotifyStatus(gpioR)
				dao.SetGpioData(gpioR)
			}
		}

		time.Sleep(time.Second)
	}
}

func AddResetData(wtime uint32, output string) {
	if !initResetTimeData {
		resetTimeData = make([]dto.ResetTimeData,0)
		initResetTimeData = true
		go ResetGpoStateFun()
	}
	
	found := false

	for i:=0;i<len(resetTimeData);i++ { 
		if resetTimeData[i].Output == output {
			resetTimeData[i].Time = wtime
			resetTimeData[i].StartTime = time.Now()
			found = true
		}
	}

	if !found {
		var lResetTimeData dto.ResetTimeData
		lResetTimeData.Time = wtime
		lResetTimeData.StartTime = time.Now()
		lResetTimeData.Output = output
		resetTimeData = append(resetTimeData, lResetTimeData)
	}
	if mutexLocked {
		resetMutex.Unlock()
		mutexLocked = false
	}
}

func GpioWrite(gpioWriteDto dto.GpioWriteDto, reset bool) bool {
	rwMutex.Lock()
	gpioData := dao.GetGpioData()
	cx := context.GetContext()
	cx.GpioDrv.Write(gpioData, gpioWriteDto)
	rwMutex.Unlock()

	rdw := cx.HwLibs.GetResetDataByWriteDto(gpioWriteDto)
	
	if !reset {
		rsMutex.Lock()
		for i:=0; i<len(rdw); i++ {
			AddResetData(rdw[i].Time, rdw[i].Output)
		}
		rsMutex.Unlock()
	}

	return true
}