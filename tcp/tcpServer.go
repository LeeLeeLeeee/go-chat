package tcpserver

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"sync"
	"time"
)

var (
	ErrorUserConnection   = errors.New("user could't connected")
	ErrorServerConnection = errors.New("fail to connect server")
)

type TcpServer struct{}

func (ts *TcpServer) ServiceServer(wg *sync.WaitGroup) {

	file, _ := os.Create("server_connect_log.txt")
	fmt.Println("Starting Server At: " + time.Now().Format("2006-01-02 15:04:05 Monday"))
	server_connection, error := net.Listen("tcp", ":2021")

	defer func() {
		file.WriteString("Close Server At: " + time.Now().Format("2006-01-02 15:04:05 Monday"))
		server_connection.Close()
		file.Close()
		os.Remove(file.Name())
		wg.Done()
	}()

	if error != nil {
		fmt.Println(ErrorServerConnection.Error())
		return
	}
	for {
		user_connection, error := server_connection.Accept()
		if error != nil {
			fmt.Println(ErrorUserConnection.Error())
			continue
		}
		go connectUserToChat(user_connection)
		file.WriteString("##Connected user: " + user_connection.RemoteAddr().String() + "\n")
	}
}

func connectUserToChat(user_connection net.Conn) {
	// var data []byte
	buffer := make([]byte, 1024)

	for {
		byteCount, error := user_connection.Read(buffer)
		if error != nil {
			if error == io.EOF {
				break
			} else {
				fmt.Println("Found Error At Client Reading:", error)
				return
			}
		}

		buffer = bytes.Trim(buffer[:byteCount], "\x00")
		// fmt.Println(string(buffer))
		// data = append(data, buffer...)
		// if data[len(data)-1] == endLineChar { //End of message, break then
		// 	break
		// }
	}

}
