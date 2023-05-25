package NKIOLCDto

type GpioNKIOLCDtoStruct struct {
	X[16] byte
	Y[16] byte
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
	Y0 byte
	Y1 byte
	Y2 byte
	Y3 byte
	Y4 byte
	Y5 byte
	Y6 byte
	Y7 byte
	Y8 byte
	Y9 byte
	Y10 byte
	Y11 byte
	Y12 byte
	Y13 byte
	Y14 byte
	Y15 byte
}

type GpioNKIOLCSetDataEl struct {
	Update	bool
	Value	byte
	Time	uint32
}

type GpioNKIOLCSetData struct {
	Y0 GpioNKIOLCSetDataEl
	Y1 GpioNKIOLCSetDataEl
	Y2 GpioNKIOLCSetDataEl
	Y3 GpioNKIOLCSetDataEl
	Y4 GpioNKIOLCSetDataEl
	Y5 GpioNKIOLCSetDataEl
	Y6 GpioNKIOLCSetDataEl
	Y7 GpioNKIOLCSetDataEl
	Y8 GpioNKIOLCSetDataEl
	Y9 GpioNKIOLCSetDataEl
	Y10 GpioNKIOLCSetDataEl
	Y11 GpioNKIOLCSetDataEl
	Y12 GpioNKIOLCSetDataEl
	Y13 GpioNKIOLCSetDataEl
	Y14 GpioNKIOLCSetDataEl
	Y15 GpioNKIOLCSetDataEl
}
