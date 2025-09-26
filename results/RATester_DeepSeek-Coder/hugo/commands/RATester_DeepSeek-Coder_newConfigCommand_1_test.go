package commands

import (
	"fmt"
	"testing"
)

func TestNewConfigCommand_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cmd := newConfigCommand()
	if cmd == nil {
		t.Errorf("newConfigCommand() = %v, want not nil", cmd)
	}

	if cmd.Name() != "config" {
		t.Errorf("newConfigCommand().Name() = %v, want 'config'", cmd.Name())
	}

	if len(cmd.Commands()) == 0 {
		t.Errorf("newConfigCommand().Commands() = %v, want not empty", cmd.Commands())
	}
}
