package dto

import "time"

type GpioDto interface {
	
}

type GpioSendDataStruct interface {

}

type GpioWriteDto interface {

}

type ResetTimeData struct {
	Time		uint32
	StartTime 	time.Time
	Output		string
}

