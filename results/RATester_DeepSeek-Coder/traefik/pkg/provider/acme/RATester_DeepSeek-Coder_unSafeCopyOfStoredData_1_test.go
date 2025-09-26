package acme

import (
	"fmt"
	"testing"
)

func TestUnSafeCopyOfStoredData_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	store := &LocalStore{
		storedData: map[string]*StoredData{
			"resolver1": {
				Account: &Account{
					Email: "test@example.com",
				},
			},
		},
	}

	copiedData := store.unSafeCopyOfStoredData()

	if len(copiedData) != 1 {
		t.Errorf("Expected 1 item in copied data, got %d", len(copiedData))
	}

	if copiedData["resolver1"].Account.Email != "test@example.com" {
		t.Errorf("Expected email 'test@example.com', got '%s'", copiedData["resolver1"].Account.Email)
	}
}
