package NKIOLC

/*
#cgo LDFLAGS: -L../../extlib -lNP6118BaseLib -lpthread -lstdc++
#cgo CFLAGS: -I../../extlib
#include "NP6118BaseLib.h"
*/
import "C"

type WdNKIOLC struct {
	Initialized bool
}

func NewWdNKIOLC() *WdNKIOLC{
	wd := new(WdNKIOLC)
	wd.Initialized = false

	return wd
}

func (wd *WdNKIOLC) Init() bool {
	res := C.NKBASE_LibraryInit()

	if res != C.NKIO_ENOERR {
		panic("Wd library not initialized")
	} else {
		wd.Initialized = true
	}
	return wd.Initialized
}

func (wd *WdNKIOLC) StartWd(time uint8) bool{
	res := C.NKWDT_Start(1, C.uchar(time))
	return (res == C.NKIO_ENOERR)
}

func (wd *WdNKIOLC) ResetWd(time uint8) bool {
	res := C.NKWDT_Reset(C.uchar(time))
	return (res == C.NKIO_ENOERR)
}

func (wd *WdNKIOLC) StopWd() bool {
	res := C.NKWDT_Stop()
	return (res == C.NKIO_ENOERR)
}