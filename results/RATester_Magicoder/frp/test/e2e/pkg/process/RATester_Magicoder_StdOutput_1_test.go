package process

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"testing"
)

func TestStdOutput_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := &Process{
		cmd:         exec.CommandContext(ctx, "ls"),
		cancel:      cancel,
		errorOutput: &bytes.Buffer{},
		stdOutput:   &bytes.Buffer{},
	}

	err := p.Start()
	if err != nil {
		t.Fatalf("Failed to start process: %v", err)
	}

	output := p.StdOutput()
	if output == "" {
		t.Errorf("Expected non-empty output, got empty string")
	}

	err = p.Stop()
	if err != nil {
		t.Fatalf("Failed to stop process: %v", err)
	}
}
