package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":4545")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer listener.Close()
	fmt.Println("Listening for messages...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			conn.Close()
			fmt.Println("Accepting error: ", err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		input := make([]byte, (1024 * 4))
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			fmt.Println("Read error: ", err)
			break
		}

		msg := string(input[0:n])
		fmt.Println("Msg: ", msg)

		response := "sh" + msg
		conn.Write([]byte(response))
	}
}
