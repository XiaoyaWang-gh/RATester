package acme

import (
	"fmt"
	"testing"
)

func TestGetAccount_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	store := &LocalStore{
		storedData: map[string]*StoredData{
			"test": {
				Account: &Account{
					Email: "test@example.com",
				},
			},
		},
	}

	account, err := store.GetAccount("test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if account.Email != "test@example.com" {
		t.Fatalf("expected email to be 'test@example.com', got %s", account.Email)
	}
}
