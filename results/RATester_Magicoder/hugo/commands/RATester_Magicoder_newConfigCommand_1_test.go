package commands

import (
	"fmt"
	"testing"
)

func TestnewConfigCommand_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cmd := newConfigCommand()
	if cmd == nil {
		t.Fatal("newConfigCommand() should not return nil")
	}

	if cmd.Name() != "config" {
		t.Errorf("Expected Name() to return 'config', got '%s'", cmd.Name())
	}

	if len(cmd.Commands()) != 1 {
		t.Errorf("Expected 1 Commands(), got %d", len(cmd.Commands()))
	}

	// Add more tests as needed
}
