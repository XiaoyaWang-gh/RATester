package process

import (
	"fmt"
	"testing"
)

func TestNewWithEnvs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := NewWithEnvs("echo", []string{"hello"}, []string{"FOO=bar"})
	if err := p.Start(); err != nil {
		t.Fatal(err)
	}
	defer p.Stop()
	if p.cmd.Path != "echo" {
		t.Errorf("cmd.Path = %s, want %s", p.cmd.Path, "echo")
	}
	if p.cmd.Args[0] != "echo" {
		t.Errorf("cmd.Args[0] = %s, want %s", p.cmd.Args[0], "echo")
	}
	if p.cmd.Args[1] != "hello" {
		t.Errorf("cmd.Args[1] = %s, want %s", p.cmd.Args[1], "hello")
	}
	if p.cmd.Env[0] != "FOO=bar" {
		t.Errorf("cmd.Env[0] = %s, want %s", p.cmd.Env[0], "FOO=bar")
	}
	if p.cmd.Dir != "" {
		t.Errorf("cmd.Dir = %s, want %s", p.cmd.Dir, "")
	}
	if p.errorOutput.String() != "" {
		t.Errorf("errorOutput = %s, want %s", p.errorOutput.String(), "")
	}
	if p.stdOutput.String() != "hello\n" {
		t.Errorf("stdOutput = %s, want %s", p.stdOutput.String(), "hello\n")
	}
}
