package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":4474")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on port 4474")
	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("Message Received:", string(message))
		newmessage := strings.ToUpper(message)
		conn.Write([]byte(newmessage + "\n"))
	}
}
