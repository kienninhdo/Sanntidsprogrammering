package main

import (
	"fmt"
	"net"
)

func UDPrecieve() {

	//create socket from address
	address, _ := net.ResolveUDPAddr("udp", ":20011")

	//listen to socket from address
	recSocket, _ := net.ListenUDP("udp", address)
	defer recSocket.Close()

	//create buffer to receive data
	buffer := make([]byte, 1024)

	for {
		//clear buffer
		for i := range buffer {
			buffer[i] = 0
		}
		n := 0

		//receive data on the socket, recAddr is where you receive from
		n, recAddr, _ := recSocket.ReadFromUDP(buffer)

		//convert to string
		receivedString := string(buffer[:n])

		fmt.Printf("%v: %s\n", recAddr, receivedString)

	}

}

func UDPsend() {

	//create socket from address
	address, _ := net.ResolveUDPAddr("udp", ":20011")

	wrSocket, _ := net.DialUDP("udp", nil, address)
	defer wrSocket.Close()

	//create message
	message := []byte("Kæser moren din")

	//
	buffer := make([]byte, 1024)

	for {
		for i := range buffer {
			buffer[i] = 0
		}

		//write data to
		n, _ := wrSocket.Write(message)

		fmt.Printf("%v", string(buffer[:n]))
	}

}

func main() {

	go UDPsend()
	go UDPrecieve()

	select {}

}