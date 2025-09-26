package version

import (
	"fmt"
	"testing"
)

func TestNewCmd_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cmd := NewCmd()

	if cmd.Name != "version" {
		t.Errorf("Expected command name to be 'version', got '%s'", cmd.Name)
	}

	if cmd.Description != `Shows the current Traefik version.` {
		t.Errorf("Expected command description to be 'Shows the current Traefik version.', got '%s'", cmd.Description)
	}

	if cmd.Run == nil {
		t.Error("Expected command run function to be non-nil")
	}
}
