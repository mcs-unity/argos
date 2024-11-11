package connector

import "testing"

func TestStartTransactionUnavailable(t *testing.T) {
	f := GetUnAvailable()
	if f.StartTransaction(10) == nil {
		t.Error("unavailable was able to start a transaction")
	}
}
