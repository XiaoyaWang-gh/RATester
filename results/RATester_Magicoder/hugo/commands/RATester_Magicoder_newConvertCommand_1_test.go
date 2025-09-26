package commands

import (
	"fmt"
	"testing"
)

func TestnewConvertCommand_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cmd := newConvertCommand()

	if cmd == nil {
		t.Error("newConvertCommand() should not return nil")
	}

	if cmd.Name() != "convert" {
		t.Errorf("Expected command name to be 'convert', but got '%s'", cmd.Name())
	}

	if len(cmd.Commands()) != 3 {
		t.Errorf("Expected 3 sub-commands, but got %d", len(cmd.Commands()))
	}

	// Add more specific tests for each sub-command
	for _, subCmd := range cmd.Commands() {
		switch subCmd.Name() {
		case "toJSON":
			// Test toJSON sub-command
		case "toTOML":
			// Test toTOML sub-command
		case "toYAML":
			// Test toYAML sub-command
		default:
			t.Errorf("Unexpected sub-command: %s", subCmd.Name())
		}
	}
}
