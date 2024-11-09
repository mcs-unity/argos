package connector

import (
	"testing"
	"time"
)

func newCharger() IConnector {
	return &Available{ConnectorState{time.Now(), AVAILABLE, -1}}
}

func TestAvailableTransaction(t *testing.T) {
	var f IConnector = newCharger()
	if err := f.StartTransaction(); err != nil {
		t.Error(err)
	}
}

func fakeCallBack(list []IConnector, id uint8) Callback {
	return func(newState IConnector) error {
		list[id] = newState
		return nil
	}
}

func TestStartTransaction(t *testing.T) {
	f := newCharger()
	if err := f.StartTransaction(); err != nil {
		t.Error(err)
	}
}

func TestGetType(t *testing.T) {
	f := newCharger()
	if f.Type() != AVAILABLE {
		t.Errorf("expected: %s got %s", AVAILABLE, f.Type())
	}
}

func TestChangeState(t *testing.T) {
	f := newCharger()
	arr := []IConnector{f}
	if err := f.ChangeState(UNAVAILABLE, fakeCallBack(arr, 0)); err != nil {
		t.Error(err)
	}

	if _, ok := arr[0].(*Unavailable); ok == false {
		t.Error("failed to alter state")
	}
}

func TestChangeWithTransaction(t *testing.T) {
	t.Error("not implemented")
}
