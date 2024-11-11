package socket

import "testing"

func MockSocket() SocketMock {
	return SocketMock{}
}

func TestClose(t *testing.T) {
	s := MockSocket()
	if err := s.Close(); err != nil {
		t.Error(err)
	}
}

func TestCloseError(t *testing.T) {
	s := MockSocket()
	s.failClose = true

	if err := s.Close(); err == nil {
		t.Error("failed to throw error")
	}
}

func TestTerminate(t *testing.T) {
	s := MockSocket()

	if err := s.Terminate(); err != nil {
		t.Error(err)
	}
}

func TestTerminateError(t *testing.T) {
	s := MockSocket()
	s.failTerminate = true

	if err := s.Terminate(); err == nil {
		t.Error("failed to throw error")
	}
}

func TestWrite(t *testing.T) {
	s := MockSocket()

	if err := s.Write([]byte("Hello world")); err != nil {
		t.Error("failed to throw error")
	}
}

func TestWriteError(t *testing.T) {
	s := MockSocket()
	s.failWrite = true

	if err := s.Write([]byte("Hello world")); err == nil {
		t.Error("failed to throw error")
	}
}

func TestConnect(t *testing.T) {
	s := MockSocket()

	if err := s.Connect([]byte("wss://localhost:3000")); err != nil {
		t.Error("failed to throw error")
	}
}

func TestConnectError(t *testing.T) {
	s := MockSocket()
	s.failConnect = true

	if err := s.Connect([]byte("Hello world")); err == nil {
		t.Error("failed to throw error")
	}
}
