package testenv

import (
	"fmt"
	"testing"
)

func TestCommand_3(t *testing.T) {
	t.Run("TestCommand", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		cmd := Command(t, "echo", "Hello, World")
		out, err := cmd.Output()
		if err != nil {
			t.Fatalf("Command execution failed: %v", err)
		}
		if string(out) != "Hello, World\n" {
			t.Errorf("Expected 'Hello, World', got '%s'", string(out))
		}
	})
}
