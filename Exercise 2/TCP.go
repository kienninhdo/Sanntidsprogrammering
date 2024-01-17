package main

import (
	"fmt"
	"net"
)

func TCPclient() {

	//address, _ := net.ResolveTCPAddr("tcp", ":33546")

	//Connect
	socket, _ := net.Dial("tcp", "127.0.0.1:33546") //skal vi dial MED IP-adressen også
	defer socket.Close()

	//Send message
	message := "Connect to 127.0.0.1:33546\000"
	socket.Write([]byte(message))

	//Read message
	buffer := make([]byte, 1024)
	n, _ := socket.Read(buffer)

	receivedString := string(buffer[:n])
	fmt.Println("Message from server:", receivedString)
}

func TCPserver() {

	socket, _ := net.Listen("tcp", "127.0.0.1:33546") //skal vi dial MED IP-adressen også
	defer socket.Close()

	fmt.Println("Server is listening on 33546")

	for {
		newClient, _ := socket.Accept()

		//HANDLE:
		defer newClient.Close()

		//Read message
		buffer := make([]byte, 1024)
		n, _ := newClient.Read(buffer)

		receivedString := string(buffer[:n])
		fmt.Println("Message from client:", receivedString)

		// //Send message
		// address, _ := net.ResolveTCPAddr("tcp", ":20011")
		// socket, _ := net.DialTCP("tcp", nil, address)
		// defer socket.Close()
		// socket.Write([]byte(receivedString))

	}
}

func main() {

	go TCPclient()
	go TCPserver()

	fmt.Printf("this is a test")

	select {}

}
