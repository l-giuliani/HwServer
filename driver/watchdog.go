package driver

type Watchdog interface {
	Init() bool
	StartWd(time uint8) bool
	ResetWd(time uint8) bool
	StopWd() bool
}