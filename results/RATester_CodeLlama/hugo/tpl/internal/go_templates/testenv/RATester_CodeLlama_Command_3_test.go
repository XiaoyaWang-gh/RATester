package testenv

import (
	"fmt"
	"testing"
)

func TestCommand_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	cmd := Command(t, "echo", "hello")
	if cmd.Path != "echo" {
		t.Errorf("cmd.Path = %q; want %q", cmd.Path, "echo")
	}
	if len(cmd.Args) != 1 || cmd.Args[0] != "echo" {
		t.Errorf("cmd.Args = %q; want %q", cmd.Args, []string{"echo"})
	}
	if cmd.Dir != "" {
		t.Errorf("cmd.Dir = %q; want %q", cmd.Dir, "")
	}
}
