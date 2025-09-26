package commands

import (
	"fmt"
	"testing"
)

func TestnewGenCommand_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cmd := newGenCommand()

	if cmd == nil {
		t.Error("newGenCommand() should not return nil")
	}

	if cmd.Name() != "gen" {
		t.Errorf("Expected command name to be 'gen', but got '%s'", cmd.Name())
	}

	commands := cmd.Commands()
	if len(commands) != 4 {
		t.Errorf("Expected 4 commands, but got %d", len(commands))
	}

	for _, c := range commands {
		if c == nil {
			t.Error("Command should not be nil")
		}
	}
}
