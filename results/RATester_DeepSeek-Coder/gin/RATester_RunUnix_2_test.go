package gin

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestRunUnix_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	engine := New()
	file := "/tmp/test.sock"

	// Remove the socket file if it exists
	if _, err := os.Stat(file); err == nil {
		if err := os.Remove(file); err != nil {
			t.Fatalf("Failed to remove existing socket file: %v", err)
		}
	}

	// Start the server in a goroutine
	go func() {
		if err := engine.RunUnix(file); err != nil {
			t.Errorf("Failed to start server: %v", err)
		}
	}()

	// Wait for the server to start
	time.Sleep(100 * time.Millisecond)

	// Check if the socket file exists
	if _, err := os.Stat(file); err != nil {
		t.Errorf("Failed to create socket file: %v", err)
	}

	// TODO: Add more tests here, for example:
	// - Check if the server is listening on the correct file
	// - Send a request to the server and check the response
	// - Check if the server correctly handles different types of requests

	// Remove the socket file
	if err := os.Remove(file); err != nil {
		t.Errorf("Failed to remove socket file: %v", err)
	}
}
