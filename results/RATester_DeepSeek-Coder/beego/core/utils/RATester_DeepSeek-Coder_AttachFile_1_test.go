package utils

import (
	"fmt"
	"net/smtp"
	"testing"
)

func TestAttachFile_1(t *testing.T) {
	e := &Email{
		Auth:     smtp.PlainAuth("", "user@example.com", "password", "mail.example.com"),
		Username: "user@example.com",
		Password: "password",
		Host:     "mail.example.com",
		Port:     587,
		From:     "from@example.com",
	}

	t.Run("Test with valid file", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := e.AttachFile("testdata/test.txt")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("Test with invalid file", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := e.AttachFile("testdata/invalid.txt")
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("Test with too many arguments", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := e.AttachFile("testdata/test.txt", "id1", "id2")
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
