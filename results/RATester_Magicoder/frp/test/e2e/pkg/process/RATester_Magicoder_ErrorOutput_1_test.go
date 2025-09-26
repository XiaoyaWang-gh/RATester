package process

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"testing"
)

func TestErrorOutput_1(t *testing.T) {
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

	output := p.ErrorOutput()
	if output != "" {
		t.Errorf("Expected empty error output, got: %s", output)
	}

	err = p.Stop()
	if err != nil {
		t.Fatalf("Failed to stop process: %v", err)
	}

	output = p.ErrorOutput()
	if output == "" {
		t.Error("Expected non-empty error output after stopping the process")
	}
}
