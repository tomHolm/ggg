package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:4545")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter a message: ")
		data, _, err := reader.ReadLine()

		if err != nil {
			fmt.Println("Read error: ", err)
			continue
		}

		if n, err := conn.Write(data); n == 0 || err != nil {
			fmt.Println("Send error: ", err)
			return
		}

		fmt.Print("Response: ")
		buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		if err != nil {
			fmt.Println("Error while read response: ", err)
			break
		}
		fmt.Print(string(buff[0:n]))
		fmt.Println()
	}
}
