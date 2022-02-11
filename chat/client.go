package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	fmt.Printf("Enter your name: ")
	var name string
	fmt.Scanln(&name)

	conn, err := net.Dial("tcp", "localhost:77")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	go handleListening(conn)
	for {
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			line := scanner.Text()
			fmt.Fprintf(conn, fmt.Sprintf("[%v] %v", name, line))
		}
	}
}

func handleListening(conn net.Conn) {
	tmp := make([]byte, 256)
	defer conn.Close()
	for {
		_, err := conn.Read(tmp)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s", tmp[:])
		tmp = make([]byte, 256)
	}
}
