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

	cmd := exec.CommandContext(ctx, "ls")
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
		t.Errorf("Expected no error, but got %v", err)
	}
}
