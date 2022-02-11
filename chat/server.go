package main

import (
	"fmt"
	"net"
)

var clients []net.Conn = make([]net.Conn, 10)

func main() {
	ln, err := net.Listen("tcp", ":77")
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		clients = append(clients, conn)
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	tmp := make([]byte, 256)
	defer conn.Close()
	for {
		_, err := conn.Read(tmp)
		if err != nil {
			fmt.Println(err)
			return
		}
		msg := fmt.Sprintf("%s\n", tmp[:])
		sendMessageToClients(conn, msg)
		fmt.Printf(msg)
		tmp = make([]byte, 256)
	}
}

func sendMessageToClients(sender net.Conn, msg string) {
	for _, conn := range clients {
		if conn != nil && conn != sender {
			conn.Write([]byte(msg))
		}
	}
}
