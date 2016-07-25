package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Print("Echo Service starting ....... ")
	service := ":2020"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	listener, err := net.ListenTCP("tcp", tcpAddr)

	if err != nil {
		fmt.Println("[ Failed ]")
		fmt.Println(err)
		os.Exit(0)
	} else {
		fmt.Println("[ Ok ]")
		fmt.Printf("Listening for connections on localhost%v\n", service)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			println(err)
		}
		fmt.Printf("Connection recived from : %s\n", conn.RemoteAddr())
		buf := make([]byte, 1024)
		inputLength, err := conn.Read(buf)
		if err != nil {
			println(err)
		}
		fmt.Printf("Message recived: %s", buf)
		conn.Write(buf[0:inputLength])
		fmt.Printf("Message echoed: %s", buf)
		conn.Close()
	}
}
