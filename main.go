package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	t "tcpgo.com/tcpserver"
)

const (
	OUT_SERVER = iota
	CONNECT_SERVER
	CREATE_ROOM
	SHOW_ROOM_LIST
	SHOW_USER_LIST
)

func main() {
	var wg sync.WaitGroup
	tcpChannel := make(chan int)
	fmt.Print("\033[H\033[2J")
	tcpServer := new(t.TcpServer)
	fmt.Println(os.Getenv("mode"))
	wg.Add(1)
	go tcpServer.ServiceServer(&wg, tcpChannel)
	time.Sleep(time.Millisecond * 3)

	var command string
	var description string
	description = "\n\n##### TCP SERVER COMMAND ####\n"
	description += "!! in the first you have to complete to connect server !!\n"
	description += "@ connect server => 1\n"
	description += "@ create root => 2\n"
	description += "@ show room list => 3\n"
	description += "@ show user list => 4\n"
	description += "@ Exit tcp server => 0\n"

	for {
		fmt.Println(description)
		fmt.Print("Please input command >>> ")
		fmt.Scanln(&command)

		switch c, _ := strconv.Atoi(command); c {
		case CONNECT_SERVER:
			tcpServer.ConnectUser(tcpChannel)
		case CREATE_ROOM:
			tcpChannel <- CREATE_ROOM
			fmt.Println("------------\ncreate room!")
		case SHOW_ROOM_LIST:
			tcpChannel <- SHOW_ROOM_LIST
			fmt.Println("------------\nshow room list")
		case SHOW_USER_LIST:
			tcpServer.GetUserList()
		case OUT_SERVER:
			tcpChannel <- OUT_SERVER
			goto Exit
		default:
			fmt.Println("------------")
		}
	}

Exit:
	wg.Wait()

}
