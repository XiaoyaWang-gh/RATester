package commands

import (
	"fmt"
	"testing"

	"github.com/bep/simplecobra"
)

func TestCommands_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	serverCommand := &serverCommand{
		commands: []simplecobra.Commander{
			// Add your test commands here
		},
	}

	expectedCommands := serverCommand.Commands()

	if len(expectedCommands) != len(serverCommand.commands) {
		t.Errorf("Expected %d commands, but got %d", len(serverCommand.commands), len(expectedCommands))
	}

	for i, command := range expectedCommands {
		if command != serverCommand.commands[i] {
			t.Errorf("Expected command at index %d to be %v, but got %v", i, serverCommand.commands[i], command)
		}
	}
}
