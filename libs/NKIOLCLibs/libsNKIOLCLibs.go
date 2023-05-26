package NKIOLCLibs

import "it.etg/gpioServer/dto"
import "it.etg/gpioServer/dto/NKIOLCDto"
import "encoding/json"

type HwLibsNKIOLC struct {

}

func NewHwLibs() *HwLibsNKIOLC {
	return new(HwLibsNKIOLC)
}

func (hwLibs *HwLibsNKIOLC) ParseWriteData(data []byte) (bool, dto.GpioWriteDto) {
	var d NKIOLCDto.GpioNKIOLCSetData
	err := json.Unmarshal(data, &d)
	if err != nil {
		return false, nil
	}
	return true, d
}

// func (hwLibs *HwLibsNKIOLC) ParseData(data []byte) (bool, dto.GpioDto) {
// 	var d NKIOLCDto.GpioNKIOLCSendData
// 	var gpioSpecDto NKIOLCDto.GpioNKIOLCDtoStruct

// 	err := json.Unmarshal(data, &d)
// 	if err != nil {
// 		return false, gpioSpecDto
// 	}

// 	gpioSpecDto.X[0] = d.X0
// 	gpioSpecDto.X[1] = d.X1 
// 	gpioSpecDto.X[2] = d.X2 
// 	gpioSpecDto.X[3] = d.X3 
// 	gpioSpecDto.X[4] = d.X4 
// 	gpioSpecDto.X[5] = d.X5 
// 	gpioSpecDto.X[6] = d.X6 
// 	gpioSpecDto.X[7] = d.X7 
// 	gpioSpecDto.X[8] = d.X8 
// 	gpioSpecDto.X[9] = d.X9 
// 	gpioSpecDto.X[10] = d.X10  
// 	gpioSpecDto.X[11] = d.X11  
// 	gpioSpecDto.X[12] = d.X12 
// 	gpioSpecDto.X[13] = d.X13 
// 	gpioSpecDto.X[14] = d.X14 
// 	gpioSpecDto.X[15] = d.X15 

// 	return true, gpioSpecDto
// }

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

	d.Y0 = gpioSpecDto.Y[0]
	d.Y1 = gpioSpecDto.Y[1]
	d.Y2 = gpioSpecDto.Y[2]
	d.Y3 = gpioSpecDto.Y[3]
	d.Y4 = gpioSpecDto.Y[4]
	d.Y5 = gpioSpecDto.Y[5]
	d.Y6 = gpioSpecDto.Y[6]
	d.Y7 = gpioSpecDto.Y[7]
	d.Y8 = gpioSpecDto.Y[8]
	d.Y9 = gpioSpecDto.Y[9]
	d.Y10 = gpioSpecDto.Y[10]
	d.Y11 = gpioSpecDto.Y[11]
	d.Y12 = gpioSpecDto.Y[12]
	d.Y13 = gpioSpecDto.Y[13]
	d.Y14 = gpioSpecDto.Y[14]
	d.Y15 = gpioSpecDto.Y[15]

	strD, _ := json.Marshal(d)
	return strD
}

func (hwLibs *HwLibsNKIOLC) CheckGpioDataChanged(oldData dto.GpioDto, newData dto.GpioDto) bool{
	oldDataDto := oldData.(NKIOLCDto.GpioNKIOLCDtoStruct)
	newDataDto := newData.(NKIOLCDto.GpioNKIOLCDtoStruct)

	for i:=0; i<16; i++ {
		if oldDataDto.X[i] != newDataDto.X[i] {
			return true
		}
	}
	for i:=0; i<16; i++ {
		if oldDataDto.Y[i] != newDataDto.Y[i] {
			return true
		}
	}
	return false
}

func (hwLibs *HwLibsNKIOLC) GetOutputDataByDo(output string) dto.GpioWriteDto {
	var gpoData NKIOLCDto.GpioNKIOLCSetData
	var gpoDataEl NKIOLCDto.GpioNKIOLCSetDataEl

	gpoDataEl.Update = true
	gpoDataEl.Value	= 0

	if output == "Y0" {
		gpoData.Y0 = gpoDataEl
	} else if output == "Y1" {
		gpoData.Y1 = gpoDataEl
	} else if output == "Y2" {
		gpoData.Y2 = gpoDataEl
	} else if output == "Y3" {
		gpoData.Y3 = gpoDataEl
	} else if output == "Y4" {
		gpoData.Y4 = gpoDataEl
	} else if output == "Y5" {
		gpoData.Y5 = gpoDataEl
	} else if output == "Y6" {
		gpoData.Y6 = gpoDataEl
	} else if output == "Y7" {
		gpoData.Y7 = gpoDataEl
	} else if output == "Y8" {
		gpoData.Y8 = gpoDataEl
	} else if output == "Y9" {
		gpoData.Y9 = gpoDataEl
	} else if output == "Y10" {
		gpoData.Y10 = gpoDataEl
	} else if output == "Y11" {
		gpoData.Y11 = gpoDataEl
	} else if output == "Y12" {
		gpoData.Y12 = gpoDataEl
	} else if output == "Y13" {
		gpoData.Y13 = gpoDataEl
	} else if output == "Y14" {
		gpoData.Y14 = gpoDataEl
	} else if output == "Y15" {
		gpoData.Y15 = gpoDataEl
	}

	return gpoData
}