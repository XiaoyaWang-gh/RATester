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
		To:      []string{"test@example.com"},
		From:    "test@example.com",
		Subject: "Test",
		Text:    "Test",
	}

	err := e.Send()
	if err != nil {
		t.Error(err)
	}
}
