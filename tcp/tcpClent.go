package tcpserver

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

const (
	OUT_SERVER = iota
	CONNECT_SERVER
	CREATE_ROOM
	SHOW_ROOM_LIST
	SOME_ONE_CHAT
)

type TcpClient struct {
	users   []*User
	usermap map[int]*User
}

const endLineChar = 10

func (tc *TcpClient) ServiceClient(wg *sync.WaitGroup) {
	tc.usermap = make(map[int]*User)
	var command string
	var description string
	description = "\n\n##### TCP SERVER COMMAND ####\n"
	description += "!! in the first you have to complete to connect server !!\n"
	description += "@ create user and connecting user to server => 1\n"
	description += "@ create root => 2\n"
	description += "@ show room list => 3\n"
	description += "@ some one chat => 4\n"
	description += "@ Exit tcp server => 0\n"

	defer func() {
		for _, user := range tc.users {
			user.connection.Close()
			// os.Remove(user.chatHistoryFile.Name())
		}
		wg.Done()
	}()

	for {
		fmt.Println(description)
		fmt.Print("Please input command >>> ")
		fmt.Scanln(&command)

		switch c, _ := strconv.Atoi(command); c {
		case CONNECT_SERVER:
			tc.ConnectUserWithCreate()
		case CREATE_ROOM:
			fmt.Println("------------\ncreate room!")
		case SHOW_ROOM_LIST:
			fmt.Println("------------\nshow room list")
		case SOME_ONE_CHAT:

		case OUT_SERVER:
			goto Exit
		default:
			fmt.Println("------------")
		}
	}
Exit:
	return
}

func (tc *TcpClient) ConnectUserWithCreate() error {
	var user_connection net.Conn
	var err error

	for {
		fmt.Println("\nConnecting To 127.0.0.1:2021")
		user_connection, err = net.Dial("tcp", ":2021")
		if err == nil {
			user, err := tc.createUser(user_connection)
			if err != nil {
				fmt.Println(err.Error())
				break
			} else {
				tc.insertUser(user)
			}
			// go tc.connectUserToChat(user_connection)
			break
		}
		time.Sleep(time.Second)
	}
	return nil
}

func (tc *TcpClient) createUser(user_connection net.Conn) (*User, error) {
	var name string
	fmt.Print("Please set your name: ")
	fmt.Scanln(&name)
	user := NewUser(len(tc.users)+1, name, user_connection)
	if user == nil {
		return nil, errors.New("fail to create user")
	}
	helloText := "Hello I'm " + name

	for _, user := range tc.users {
		_, err := user.connection.Write([]byte(helloText))
		if err != nil {
			fmt.Println(err)
		}
	}
	return user, nil
}

func (tc *TcpClient) insertUser(user *User) {
	tc.users = append(tc.users, user)
	tc.usermap[user.id] = user
}

// func (tc *TcpClient) connectUserToChat(user_connection net.Conn) {
// 	// var data []byte
// 	buffer := make([]byte, 1024)

// 	for {
// 		byteCount, error := user_connection.Read(buffer)
// 		if error != nil {
// 			if error == io.EOF {
// 				break
// 			} else {
// 				fmt.Println("Found Error At Client Reading:", error)
// 				return
// 			}
// 		}

// 		buffer = bytes.Trim(buffer[:byteCount], "\x00")
// 		// fmt.Println(string(buffer))
// 		// data = append(data, buffer...)
// 		// if data[len(data)-1] == endLineChar { //End of message, break then
// 		// 	break
// 		// }
// 	}

// }
