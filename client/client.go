package main

import (
	"errors"
	"fmt"
	"net"
)

var (
	dest  string = "localhost:8080"
	proto string = "tcp"
)

func main() {
	fmt.Println("Start...")

	conn, err := net.Dial(proto, dest)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Connected")

	for {
		var inputStr string
		fmt.Println("Please enter your input:")
		fmt.Scanf("%s", &inputStr)

		if inputStr == "Break" {
			break
		}

		_, err = conn.Write([]byte(inputStr))
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Write")
		output, err := handleConnection(conn)
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Printf("Your result is:\n%v\n\n", output)
	}

	conn.Close()
}

func handleConnection(conn net.Conn) (string, error) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("cannot read")
	}

	return string(buf), nil
}

// tcp, udp
