package NKIOLCDto

import "encoding/json"

type GpioNKIOLCDtoStruct struct {
	X[16] byte
}

type GpioNKIOLCSendData struct {
	X0 byte
	X1 byte
	X2 byte
	X3 byte
	X4 byte
	X5 byte
	X6 byte
	X7 byte
	X8 byte
	X9 byte
	X10 byte
	X11 byte
	X12 byte
	X13 byte
	X14 byte
	X15 byte
}

func (gpioVal GpioNKIOLCDtoStruct) GetFormattedData() []byte {
	d := GpioNKIOLCSendData{}
	d.X0 = gpioVal.X[0]
	d.X1 = gpioVal.X[1]
	d.X2 = gpioVal.X[2]
	d.X3 = gpioVal.X[3]
	d.X4 = gpioVal.X[4]
	d.X5 = gpioVal.X[5]
	d.X6 = gpioVal.X[6]
	d.X7 = gpioVal.X[7]
	d.X8 = gpioVal.X[8]
	d.X9 = gpioVal.X[9]
	d.X10 = gpioVal.X[10]
	d.X11 = gpioVal.X[11]
	d.X12 = gpioVal.X[12]
	d.X13 = gpioVal.X[13]
	d.X14 = gpioVal.X[14]
	d.X15 = gpioVal.X[15]

	strD, _ := json.Marshal(d)
	return strD
}