package utils

import (
	"fmt"
	"testing"
)

func TestSend_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &Email{
		Username: "test@example.com",
		Password: "password",
		Host:     "smtp.example.com",
		Port:     587,
		To:       []string{"recipient@example.com"},
	}

	err := e.Send()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
