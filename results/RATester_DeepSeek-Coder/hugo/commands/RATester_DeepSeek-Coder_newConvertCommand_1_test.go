package commands

import (
	"fmt"
	"testing"
)

func TestNewConvertCommand_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cmd := newConvertCommand()

	if cmd == nil {
		t.Errorf("newConvertCommand() returned nil")
	}

	if cmd.Name() != "convert" {
		t.Errorf("newConvertCommand() returned command with wrong name, expected 'convert', got %s", cmd.Name())
	}

	if len(cmd.Commands()) != 3 {
		t.Errorf("newConvertCommand() returned command with wrong number of sub-commands, expected 3, got %d", len(cmd.Commands()))
	}

	// Add more specific tests for the sub-commands here
}
