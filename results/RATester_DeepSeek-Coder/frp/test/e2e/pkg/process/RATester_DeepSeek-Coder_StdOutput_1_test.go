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

	cmd := exec.CommandContext(ctx, "echo", "test")
	stdOutput := &bytes.Buffer{}
	errorOutput := &bytes.Buffer{}
	cmd.Stdout = stdOutput
	cmd.Stderr = errorOutput

	p := &Process{
		cmd:         cmd,
		cancel:      cancel,
		errorOutput: errorOutput,
		stdOutput:   stdOutput,
	}

	err := p.Start()
	if err != nil {
		t.Fatalf("Failed to start process: %v", err)
	}

	output := p.StdOutput()
	if output != "test\n" {
		t.Errorf("Expected output 'test\\n', got '%s'", output)
	}

	err = p.Stop()
	if err != nil {
		t.Fatalf("Failed to stop process: %v", err)
	}
}
