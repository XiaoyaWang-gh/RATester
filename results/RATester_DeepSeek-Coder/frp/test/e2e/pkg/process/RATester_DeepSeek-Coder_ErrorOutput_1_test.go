package process

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"testing"
	"time"
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
		cmd:         exec.CommandContext(ctx, "echo", "test"),
		cancel:      cancel,
		errorOutput: &bytes.Buffer{},
		stdOutput:   &bytes.Buffer{},
	}

	err := p.Start()
	if err != nil {
		t.Fatalf("Failed to start process: %v", err)
	}

	time.Sleep(100 * time.Millisecond)

	p.Stop()

	expectedOutput := "test\n"
	if p.ErrorOutput() != expectedOutput {
		t.Errorf("Expected error output to be '%s', but got '%s'", expectedOutput, p.ErrorOutput())
	}
}
