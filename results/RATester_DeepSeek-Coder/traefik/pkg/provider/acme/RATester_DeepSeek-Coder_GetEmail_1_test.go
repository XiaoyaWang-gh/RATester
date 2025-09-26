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

	acc := &Account{
		Email: "test@example.com",
	}

	email := acc.GetEmail()

	if email != "test@example.com" {
		t.Errorf("Expected email to be 'test@example.com', got '%s'", email)
	}
}
