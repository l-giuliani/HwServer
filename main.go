package main

import "fmt"
import "it.etg/gpioServer/services"
import "it.etg/gpioServer/server"
import "it.etg/gpioServer/libs/network"

func main () {
	fmt.Println("test")
	udpClient := network.NewUdpClient()
	udpClient.WriteAndClose("192.168.55.146:5683", []byte("Test Data !!!"))
	services.SystemInit()
	server.Init()
}