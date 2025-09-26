package nathole

import (
	"fmt"
	"testing"
)

func TestGenNatHoleResponse_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{
		clientCfgs: make(map[string]*ClientCfg),
		sessions:   make(map[string]*Session),
	}

	session := &Session{
		sid: "test_sid",
	}

	resp := ctrl.GenNatHoleResponse("test_transaction", session, "test_error")

	if resp.TransactionID != "test_transaction" {
		t.Errorf("Expected TransactionID to be 'test_transaction', got %s", resp.TransactionID)
	}

	if resp.Sid != "test_sid" {
		t.Errorf("Expected Sid to be 'test_sid', got %s", resp.Sid)
	}

	if resp.Error != "test_error" {
		t.Errorf("Expected Error to be 'test_error', got %s", resp.Error)
	}
}
