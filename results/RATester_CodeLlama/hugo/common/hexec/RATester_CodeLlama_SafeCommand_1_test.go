package hexec

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSafeCommand_1(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()

		name := "ls"
		arg := []string{"-l", "-a"}

		cmd, err := SafeCommand(name, arg...)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if cmd.Path != name {
			t.Errorf("unexpected cmd.Path: %v", cmd.Path)
		}

		if !reflect.DeepEqual(cmd.Args, arg) {
			t.Errorf("unexpected cmd.Args: %v", cmd.Args)
		}
	})

	t.Run("failure", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()

		name := "not-exist"
		arg := []string{"-l", "-a"}

		cmd, err := SafeCommand(name, arg...)
		if err == nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if cmd != nil {
			t.Errorf("unexpected cmd: %v", cmd)
		}
	})
}
