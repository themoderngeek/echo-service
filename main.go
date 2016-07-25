package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	fmt.Print("Echo Service starting ....... ")
	service := ":1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	listener, err := net.ListenTCP("tcp", tcpAddr)

	if err != nil {
		fmt.Println("[ Failed ]")
		os.Exit(0)
	} else {
		fmt.Println("[ Ok ]")
		fmt.Printf("Listening for connections on localhost%v\n", service)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {

		}
		fmt.Printf("Connection recived from : %s\n", conn.RemoteAddr())
		buf := make([]byte, 1024)
		inputLength, err := conn.Read(buf)
		if err != nil {
			conn.Close()
		}
		fmt.Printf("Message recived: %s", buf)
		conn.Write(buf[0:inputLength])
		fmt.Printf("Message echoed: %s", buf)
		conn.Close()
	}
}
