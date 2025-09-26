package acme

import (
	"fmt"
	"testing"
)

func TestGetEmail_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	account := &Account{
		Email: "test@example.com",
	}

	email := account.GetEmail()

	if email != "test@example.com" {
		t.Errorf("Expected email to be 'test@example.com', but got '%s'", email)
	}
}
