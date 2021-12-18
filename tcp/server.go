package tcpserver

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	ErrorUserConnection   = errors.New("user could't connected")
	ErrorServerConnection = errors.New("fail to connect server")
)

type TcpInterface interface {
	ServicesServer(sync.WaitGroup, chan int)
	ConnectUser(chan int) net.Conn
	insertUser(User)
}

type TcpServer struct {
	users   []User
	usermap map[int]User
}

const (
	OUT_SERVER = iota
	CONNECT_SERVER
	CREATE_ROOM
	SHOW_ROOM_LIST
	SHOW_USER_LIST
)

func (ts *TcpServer) ServiceServer(wg *sync.WaitGroup, ch chan int) {
	ts.usermap = make(map[int]User)
	file, _ := os.Create("server_connect_log.txt")
	fmt.Println("Starting Server At: " + time.Now().Format("2006-01-02 15:04:05 Monday"))
	server_connection, error := net.Listen("tcp", ":2021")

	defer func() {
		file.WriteString("Close Server At: " + time.Now().Format("2006-01-02 15:04:05 Monday"))
		server_connection.Close()
		file.Close()
		os.Remove(file.Name())
		close(ch)
		wg.Done()
	}()

	if error != nil {
		fmt.Println(ErrorServerConnection.Error())
		return
	}

	for mode := range ch {
		switch mode {
		case OUT_SERVER:
			return
		case CONNECT_SERVER:
			user_connection, error := server_connection.Accept()
			if error != nil {
				fmt.Println(ErrorUserConnection.Error())
				continue
			}
			file.WriteString("##Connected user: " + user_connection.RemoteAddr().String() + "\n")
		default:
			continue
		}
	}

}

func (ts *TcpServer) ConnectUser(ch chan int) {
	var user_connection net.Conn
	var err error

	for {
		fmt.Println("\nConnecting To 127.0.0.1:2021")
		user_connection, err = net.Dial("tcp", ":2021")
		if err == nil {
			var name string
			fmt.Print("Please set your name: ")
			fmt.Scanln(&name)
			user := User{id: len(ts.users) + 1, netInfo: user_connection, name: name}
			ts.insertUser(user)
			ch <- CONNECT_SERVER
			break
		} else {
			fmt.Print("*")
		}
		time.Sleep(time.Second)
	}
}

func (ts *TcpServer) insertUser(user User) {
	ts.users = append(ts.users, user)
	ts.usermap[user.id] = user
}

func (ts *TcpServer) GetUserList() {
	var s []string
	for _, e := range ts.users {
		s = append(s, e.GetName())
	}
	fmt.Println(strings.Join(s, ", "))
}
