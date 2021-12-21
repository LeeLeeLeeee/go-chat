package tcpserver

import (
	"errors"
	"net"
	"os"
)

type UserInterface interface {
	CreateRoom() (int, error)
	EnterRoom() (bool, error)
	SendMessage() error
	ExitRoom() error
	ExitChatServer() error
	GetName() string
}

type AdminInterface interface {
	UserInterface
	DeleteRoom() error
}

type User struct {
	id              int
	name            string
	connection      net.Conn
	chatHistoryFile *os.File
}

func NewUser(id int, name string, user_connection net.Conn) *User {
	// file, err := os.Create(name + ".txt")
	// if err != nil {
	// 	return nil
	// }
	return &User{id: id, name: name, connection: user_connection, chatHistoryFile: nil}
}

func (user *User) CreateRoom() (int, error) {
	return 0, nil
}

func (user *User) EnterRoom() (bool, error) {
	return true, nil
}

func (user *User) SendMessage() (string, error) {
	return "", nil
}

func (user *User) ExitRoom() error {
	return nil
}

func (user *User) GetName() string {
	return user.name
}

func (user *User) ExitChatServer() error {
	err := user.connection.Close()
	if err != nil {
		return errors.New("exit chat server fail")
	}
	return nil
}
