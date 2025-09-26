package process

import (
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

	p := &Process{
		cmd: &exec.Cmd{
			Path: "echo",
			Args: []string{"echo", "hello"},
		},
	}
	err := p.Start()
	if err != nil {
		t.Fatal(err)
	}
	if p.cmd.Process == nil {
		t.Fatal("process is not started")
	}
}
