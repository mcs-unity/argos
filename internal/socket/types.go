package socket

import "github.com/gorilla/websocket"

type ISocket interface {
	Connect(url []byte) error
	Terminate() error
	Read()
	Write(data []byte) error
	Close() error
}

type Socket struct {
	con *websocket.Conn
}

type SocketMock struct {
	failConnect   bool
	failTerminate bool
	failClose     bool
	failWrite     bool
}
