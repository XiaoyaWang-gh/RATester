package plugins

import (
	"fmt"
	"testing"
)

func TestResetAll_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{
		goPath: "/path/to/go",
	}

	err := client.ResetAll()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
