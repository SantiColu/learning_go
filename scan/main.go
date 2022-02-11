package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

var address string = "localhost"

func main() {
	if len(os.Args) > 1 {
		address = os.Args[1]
	}

	fmt.Printf("Scanning ports of: %s ...\n", address)

	wg := 0
	for i := 0; i < 65535; i++ {
		go scanPort(i, &wg)
	}

	for wg > 0 {
		fmt.Printf("\rRemaining: %d                  ", wg)
	}
	fmt.Println("\rRemaining: 0                     ")

	fmt.Println("\nPress enter to exit")
	fmt.Scanln()
}

func scanPort(port int, wg *int) {
	*wg++
	startTime := time.Now()
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%v", address, port))
	if err != nil {
		*wg--
		return
	}
	conn.Close()
	fmt.Printf("\r[i] Port %v is Open! [%v]\n", port, time.Since(startTime))
	*wg--
}
