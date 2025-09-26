package commands

import (
	"context"
	"fmt"
	"testing"
)

func TestnewImportCommand_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cmd := newImportCommand()

	if cmd == nil {
		t.Error("newImportCommand() should not return nil")
	}

	if cmd.Name() != "import" {
		t.Errorf("newImportCommand().Name() = %s; want 'import'", cmd.Name())
	}

	if len(cmd.Commands()) != 1 {
		t.Errorf("newImportCommand().Commands() = %d; want 1", len(cmd.Commands()))
	}

	if err := cmd.Init(nil); err != nil {
		t.Errorf("newImportCommand().Init() returned an error: %v", err)
	}

	if err := cmd.PreRun(nil, nil); err != nil {
		t.Errorf("newImportCommand().PreRun() returned an error: %v", err)
	}

	if err := cmd.Run(context.Background(), nil, nil); err != nil {
		t.Errorf("newImportCommand().Run() returned an error: %v", err)
	}
}
