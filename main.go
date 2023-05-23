package main

import "fmt"
import "it.etg/gpioServer/services"
import "it.etg/gpioServer/dao"

func main () {
	fmt.Println("test")
	
	services.GpioRead()
	gd := dao.GetGpioData()
	fmt.Println(gd)
}