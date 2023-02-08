package main

import (
	"fmt"
	"net"
	"strings"
)

const SHORT_PREFIX = "sh"
const LONG_PREFIX = "shp"

type Reduplicator struct {
	input []byte
}

func (r Reduplicator) run() []byte {
	inputStr := string(r.input)
	words := strings.Split(inputStr, " ")
	for i := 0; i < len(words); i++ {
		words[i] = r.processWord(words[i])
	}
	return []byte(strings.Join(words, " "))
}

func (r Reduplicator) processWord(s string) string {
	s = strings.ToLower(s)
	pos := strings.IndexAny(s, "aeyiou")
	switch pos {
	case 0:
		return LONG_PREFIX + s
	case 1:
		return SHORT_PREFIX + s
	default:
		return SHORT_PREFIX + s[pos-1:]
	}
}

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

		reduplicator := Reduplicator{input}
		result := reduplicator.run()
		fmt.Println("Response: ", string(result))
		conn.Write(result)
	}
}
