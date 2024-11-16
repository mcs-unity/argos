package socket

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

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
	// todo input callback that will trigger event accepting read message payload
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
