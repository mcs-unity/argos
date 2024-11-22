package ocpp

import (
	"testing"
)

type Mock struct{ test string }

func TestGetResponse(t *testing.T) {
	uuid := "123"
	mock := Mock{"Hello"}
	p := Response(uuid, mock)

	if len(p) != 3 {
		t.Error("invalid ocpp frame length")
	}

	if p[0] != RESPONSE {
		t.Errorf("expected %f got %f", RESPONSE, p[0])
	}

	if p[1] != uuid {
		t.Errorf("expected %s got %s", uuid, p[1])
	}

	m, ok := p[2].(Mock)
	if !ok {
		t.Error("payload is not of type Mock")
	}

	if m.test != mock.test {
		t.Errorf("expected %s got %s", m.test, mock.test)
	}
}

func TestGetRequest(t *testing.T) {
	uuid := "123"
	mock := Mock{"Hello"}
	action := Action("Test")
	p := Request(uuid, action, mock)

	if len(p) != 4 {
		t.Error("invalid ocpp frame length")
	}

	if p[0] != CALL {
		t.Errorf("expected %f got %f", CALL, p[0])
	}

	if p[1] != uuid {
		t.Errorf("expected %s got %s", uuid, p[1])
	}

	if p[2] != action {
		t.Errorf("expected %s got %s", action, p[2])
	}

	m, ok := p[3].(Mock)
	if !ok {
		t.Error("payload is not of type Mock")
	}

	if m.test != mock.test {
		t.Errorf("expected %s got %s", m.test, mock.test)
	}
}
