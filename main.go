package main

import "fmt"
import "it.etg/gpioServer/services"
import "it.etg/gpioServer/server"

func main () {
	fmt.Println("test")
	services.SystemInit()
	server.Init()
}