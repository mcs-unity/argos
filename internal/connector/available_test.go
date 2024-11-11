package connector

import (
	"testing"
)

func TestGetAvailableState(t *testing.T) {
	state := GetAvailable()
	if state == nil {
		t.Error("failed to get connector state available")
	}
}

func TestAvailableTransaction(t *testing.T) {
	var f IConnector = GetAvailable()
	if err := f.StartTransaction(1); err != nil {
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
	a := GetAvailable()
	transaction := 100
	a.StartTransaction(transaction)
	if tr := a.GetTransaction(); tr != transaction {
		t.Errorf("expected transaction id %d got %d", transaction, tr)
	}
}

func TestGetType(t *testing.T) {
	f := GetAvailable()
	if f.Type() != AVAILABLE {
		t.Errorf("expected: %s got %s", AVAILABLE, f.Type())
	}
}

func TestChangeState(t *testing.T) {
	f := GetAvailable()
	arr := []IConnector{f}
	if err := f.ChangeState(UNAVAILABLE, fakeCallBack(arr, 0)); err != nil {
		t.Error(err)
	}

	if _, ok := arr[0].(*Unavailable); ok == false {
		t.Error("failed to alter state")
	}
}

func TestBadChangeState(t *testing.T) {
	f := GetAvailable()
	arr := []IConnector{f}
	if err := f.ChangeState(FINISHING, fakeCallBack(arr, 0)); err == nil {
		t.Error("state change should not be possible")
	}
}
