package main

import (
	"fmt"
	"os"
	"sync"

	t "tcpgo.com/tcpserver"
)

func main() {
	var wg sync.WaitGroup
	fmt.Print("\033[H\033[2J")
	tcpServer := new(t.TcpServer)
	fmt.Println(os.Getenv("mode"))
	wg.Add(1)
	go tcpServer.ServiceServer(&wg)
	wg.Wait()
}
