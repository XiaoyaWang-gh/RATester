package utils

import (
	"fmt"
	"testing"
)

func TestNewEMail_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	config := `{
		"identity": "test",
		"username": "test",
		"password": "test",
		"host": "test",
		"port": 25,
		"from": "test",
		"to": ["test"],
		"bcc": ["test"],
		"cc": ["test"],
		"subject": "test",
		"text": "test",
		"html": "test",
		"headers": {
			"test": "test"
		},
		"attachments": [
			{
				"reader": "test",
				"filename": "test",
				"args": ["test"]
			}
		],
		"read_receipt": ["test"]
	}`
	// Act
	e := NewEMail(config)
	// Assert
	if e == nil {
		t.Errorf("NewEMail() failed")
	}
}
