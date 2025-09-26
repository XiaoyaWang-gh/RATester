package process

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"testing"
)

func TestStart_2(t *testing.T) {
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
		t.Errorf("Expected no error, got %v", err)
	}

	if !cmd.ProcessState.Success() {
		t.Errorf("Expected process to be successful, got %v", cmd.ProcessState)
	}

	output := p.StdOutput()
	if output != "test\n" {
		t.Errorf("Expected output to be 'test\\n', got %v", output)
	}

	errOutput := p.ErrorOutput()
	if errOutput != "" {
		t.Errorf("Expected error output to be '', got %v", errOutput)
	}
}
