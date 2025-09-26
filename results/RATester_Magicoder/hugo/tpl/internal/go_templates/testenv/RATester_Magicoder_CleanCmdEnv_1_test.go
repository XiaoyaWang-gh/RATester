package testenv

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestCleanCmdEnv_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cmd := exec.Command("ls")
	cmd.Env = []string{"GODEBUG=debug", "GOTRACEBACK=traceback"}
	cmd = CleanCmdEnv(cmd)

	if len(cmd.Env) != 0 {
		t.Errorf("Expected env to be empty, but got %v", cmd.Env)
	}
}
