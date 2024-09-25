package main

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

var (
	dest  string = "localhost:8080"
	proto string = "tcp"
)

var clientNumber int = 1
var clientData map[string]string

func setup() {
	clientData = make(map[string]string, 10)
}

func main() {
	setup()

	ln, err := net.Listen(proto, dest)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Server is up in %v protocol and %v\n\n", proto, dest)

	for {
		conn, err := ln.Accept() // wait
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Client no %v was accepted\n\n", clientNumber)

		go handleClient(conn) // free
	}
}

func handleClient(conn net.Conn) {
	var clientID int = clientNumber
	clientNumber++

	for {
		output, err := handleConnection(conn)
		if err != nil {
			fmt.Println(err)
			return
		}

		outputSlc := strings.Split(output, "-")

		fmt.Printf("Your data from client no %v is:\n%v\n", clientID, output)
		fmt.Printf("Your slice data from client no %v is:\n%v\n", clientID, outputSlc)

		cmd := outputSlc[0]
		switch cmd {
		case "signUp":
			username := outputSlc[1]
			password := outputSlc[2]

			_, isSuccess := clientData[username]
			if isSuccess {
				conn.Write([]byte("Your username was signUp"))
				continue
			}

			clientData[username] = password

			conn.Write([]byte("SignUp success"))

		case "login":
			username := outputSlc[1]
			inputPassword := outputSlc[2]

			password, isSuccess := clientData[username]
			if !isSuccess {
				conn.Write([]byte("Your username was not found"))
				continue
			}

			if inputPassword != password {
				conn.Write([]byte("Your password was not match"))
				continue
			}

			conn.Write([]byte("Login success"))

		default:
			conn.Write([]byte("Your cmd is invalid"))
		}

		fmt.Println()
	}
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
