package socket

import "github.com/gorilla/websocket"

type ISocket interface {
	Connect(url []byte) error
	Terminate() error
	Close() error
}

type Socket struct {
	con *websocket.Conn
}
