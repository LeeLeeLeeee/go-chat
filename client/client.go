package main

import (
	"sync"

	t "tcpgo.com/tcpserver"
)

func main() {
	var wg sync.WaitGroup

	defer func() {
		// END..
	}()
	tcpClient := new(t.TcpClient)
	wg.Add(1)
	tcpClient.ServiceClient(&wg)
	wg.Wait()

}
