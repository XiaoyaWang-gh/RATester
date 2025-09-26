package commands

import (
	"fmt"
	"testing"
)

func TestNewImportCommand_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cmd := newImportCommand()

	if cmd == nil {
		t.Errorf("newImportCommand() returned nil")
	}

	if cmd.Name() != "import" {
		t.Errorf("newImportCommand() returned command with wrong name, expected 'import', got %s", cmd.Name())
	}

	if cmd.r != nil {
		t.Errorf("newImportCommand() returned command with non-nil rootCommand, expected nil")
	}

	if cmd.force != false {
		t.Errorf("newImportCommand() returned command with force set to true, expected false")
	}

	if len(cmd.commands) != 1 {
		t.Errorf("newImportCommand() returned command with wrong number of sub-commands, expected 1, got %d", len(cmd.commands))
	}
}
