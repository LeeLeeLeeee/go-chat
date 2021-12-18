package tcpserver

import "net"

type UserInterface interface {
	CreateRoom() (int, error)
	EnterRoom() (bool, error)
	SendMessage() error
	ExitRoom() error
	GetIp() net.Addr
}

type AdminInterface interface {
	UserInterface
	DeleteRoom() error
}

type User struct {
	id      int
	name    string
	netInfo net.Conn
}

func (user User) CreateRoom() (int, error) {
	return 0, nil
}

func (user User) EnterRoom() (bool, error) {
	return true, nil
}

func (user User) SendMessage() error {
	return nil
}

func (user User) ExitRoom() error {
	return nil
}

func (user User) GetIp() net.Addr {
	return user.netInfo.LocalAddr()
}

func (user User) GetName() string {
	return user.name
}
