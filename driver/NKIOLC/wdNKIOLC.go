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
	
}