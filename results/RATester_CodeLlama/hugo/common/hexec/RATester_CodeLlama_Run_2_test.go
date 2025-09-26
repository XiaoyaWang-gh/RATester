package hexec

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"
)

func TestRun_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &cmdWrapper{
		name: "npm",
		c: &exec.Cmd{
			Args: []string{"npm", "install", "--save-dev", "typescript"},
		},
		outerr: &bytes.Buffer{},
	}
	err := c.Run()
	if err != nil {
		t.Errorf("failed to execute binary %q with args %v: %s", c.name, c.c.Args[1:], c.outerr.String())
	}
}
