package testenv

import (
	"fmt"
	"os/exec"
	"reflect"
	"testing"
)

func TestCleanCmdEnv_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	cmd := &exec.Cmd{}
	cmd.Env = []string{"GODEBUG=foo", "GOTRACEBACK=bar", "FOO=bar"}
	CleanCmdEnv(cmd)
	if got, want := cmd.Env, []string{"FOO=bar"}; !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
