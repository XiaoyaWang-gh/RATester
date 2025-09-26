package utils

import (
	"fmt"
	"testing"
)

func TestBytes_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &Email{
		To:      []string{"test@example.com"},
		From:    "test@example.com",
		Subject: "Test Email",
	}

	_, err := e.Bytes()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
