package network

import "net"

type UdpClient struct {
	Conn *net.UDPConn 
}

func NewUdpClient() *UdpClient {
	return new(UdpClient)
}

func (udpClient *UdpClient) WriteAndClose(addressAndPort string, data[]byte) bool {
	udpAddr, err := net.ResolveUDPAddr("udp", addressAndPort)
	if err != nil {
		return false
	}
	conn, err := net.DialUDP("udp", nil, udpAddr)
	defer conn.Close()
	if err != nil {
		return false
	}
	_, err2 := conn.Write(data)
	return (err2 != nil)
}