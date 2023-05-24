package NKIOLCLibs

import "it.etg/gpioServer/dto"
import "it.etg/gpioServer/dto/NKIOLCDto"
import "encoding/json"

type HwLibsNKIOLC struct {

}

func NewHwLibs() *HwLibsNKIOLC {
	return new(HwLibsNKIOLC)
}

func (hwLibs *HwLibsNKIOLC) ParseData(data []byte) (bool, dto.GpioDto) {
	var d NKIOLCDto.GpioNKIOLCSendData
	var gpioSpecDto NKIOLCDto.GpioNKIOLCDtoStruct

	err := json.Unmarshal(data, &d)
	if err != nil {
		return false, gpioSpecDto
	}

	gpioSpecDto.X[0] = d.X0
	gpioSpecDto.X[1] = d.X1 
	gpioSpecDto.X[2] = d.X2 
	gpioSpecDto.X[3] = d.X3 
	gpioSpecDto.X[4] = d.X4 
	gpioSpecDto.X[5] = d.X5 
	gpioSpecDto.X[6] = d.X6 
	gpioSpecDto.X[7] = d.X7 
	gpioSpecDto.X[8] = d.X8 
	gpioSpecDto.X[9] = d.X9 
	gpioSpecDto.X[10] = d.X10  
	gpioSpecDto.X[11] = d.X11  
	gpioSpecDto.X[12] = d.X12 
	gpioSpecDto.X[13] = d.X13 
	gpioSpecDto.X[14] = d.X14 
	gpioSpecDto.X[15] = d.X15 

	return true, gpioSpecDto
}

func (hwLibs *HwLibsNKIOLC) GetFormattedData(gpioDto dto.GpioDto) []byte{
	d := NKIOLCDto.GpioNKIOLCSendData{}
	gpioSpecDto := gpioDto.(NKIOLCDto.GpioNKIOLCDtoStruct)

	d.X0 = gpioSpecDto.X[0]
	d.X1 = gpioSpecDto.X[1]
	d.X2 = gpioSpecDto.X[2]
	d.X3 = gpioSpecDto.X[3]
	d.X4 = gpioSpecDto.X[4]
	d.X5 = gpioSpecDto.X[5]
	d.X6 = gpioSpecDto.X[6]
	d.X7 = gpioSpecDto.X[7]
	d.X8 = gpioSpecDto.X[8]
	d.X9 = gpioSpecDto.X[9]
	d.X10 = gpioSpecDto.X[10]
	d.X11 = gpioSpecDto.X[11]
	d.X12 = gpioSpecDto.X[12]
	d.X13 = gpioSpecDto.X[13]
	d.X14 = gpioSpecDto.X[14]
	d.X15 = gpioSpecDto.X[15]

	strD, _ := json.Marshal(d)
	return strD
}