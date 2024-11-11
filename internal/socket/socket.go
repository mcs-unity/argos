package socket

import (
	"errors"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

/*
	contains websocket logic processing
	handle disconnects and shutdowns
*/

func (s *Socket) Connect(url []byte) error {
	// set read write buffer for more control
	dialer := websocket.Dialer{
		Subprotocols:      []string{"ocpp1.6"},
		EnableCompression: true,
		HandshakeTimeout:  10 * time.Second,
	}
	con, _, err := dialer.Dial(string(url), nil)
	if err != nil {
		return err
	}

	s.con = con
	return nil
}

func (s *Socket) Read() {
	for {
		_, m, err := s.con.ReadMessage()
		if err != nil {

		}

		fmt.Println(string(m))
	}
}

func (s *Socket) Write(data []byte) error {
	return s.con.WriteMessage(websocket.TextMessage, data)
}

func (s *Socket) Terminate() error {
	return s.con.Close()
}

func (s *Socket) Close() error {
	return s.con.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(
			websocket.CloseGoingAway,
			"Normal closure",
		),
		time.Now().Add(1*time.Second))
}

func (s *SocketMock) Connect(url []byte) error {
	if s.failConnect {
		return errors.New("")
	}

	return nil
}

func (s *SocketMock) Terminate() error {
	if s.failTerminate {
		return errors.New("failed to terminate socket")
	}

	return nil
}

func (s *SocketMock) Read() {

}

func (s *SocketMock) Write(data []byte) error {
	if s.failWrite {
		return errors.New("failed to send message ")
	}

	return nil
}

func (s *SocketMock) Close() error {
	if s.failClose {
		return errors.New("failed to send close message ")
	}

	return nil
}
