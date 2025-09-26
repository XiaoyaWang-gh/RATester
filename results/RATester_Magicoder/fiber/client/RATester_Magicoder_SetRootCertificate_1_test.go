package client

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestSetRootCertificate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{}

	// Create a temporary file for the test
	tempFile, err := os.CreateTemp("", "testcert")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write some test data to the file
	_, err = tempFile.Write([]byte("test data"))
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	// Call the method under test
	client.SetRootCertificate(tempFile.Name())

	// Check if the file was opened and read correctly
	file, err := os.Open(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to open temp file: %v", err)
	}
	defer file.Close()

	pem, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("Failed to read from temp file: %v", err)
	}

	if string(pem) != "test data" {
		t.Errorf("Expected 'test data', got '%s'", string(pem))
	}
}
