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
		From:    "test@example.com",
		To:      []string{"test1@example.com", "test2@example.com"},
		Subject: "Test Subject",
	}

	_, err := e.Bytes()
	if err != nil {
		t.Errorf("Failed to generate bytes: %s", err)
	}
}
