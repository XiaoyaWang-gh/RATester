package hexec

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"
)

func TestStdinPipe_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cmd := &cmdWrapper{
		name:   "test",
		c:      exec.Command("echo", "hello"),
		outerr: &bytes.Buffer{},
	}

	pipe, err := cmd.StdinPipe()
	if err != nil {
		t.Fatalf("Failed to create StdinPipe: %v", err)
	}

	_, err = pipe.Write([]byte("test"))
	if err != nil {
		t.Fatalf("Failed to write to pipe: %v", err)
	}

	err = pipe.Close()
	if err != nil {
		t.Fatalf("Failed to close pipe: %v", err)
	}

	err = cmd.Run()
	if err != nil {
		t.Fatalf("Failed to run command: %v", err)
	}

	output := cmd.outerr.String()
	if output != "test\n" {
		t.Fatalf("Unexpected output: %s", output)
	}
}
