package socket

import "errors"

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

func (s *SocketMock) Read() {}

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
