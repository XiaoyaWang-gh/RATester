package testenv

import (
	"fmt"
	"os/exec"
	"testing"

	"gotest.tools/assert"
)

func TestCleanCmdEnv_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		cmd      *exec.Cmd
		expected []string
	}{
		{
			name: "Test with empty environment",
			cmd: &exec.Cmd{
				Path: "/bin/echo",
				Args: []string{"echo", "hello", "world"},
				Env:  []string{},
			},
			expected: []string{},
		},
		{
			name: "Test with non-empty environment",
			cmd: &exec.Cmd{
				Path: "/bin/echo",
				Args: []string{"echo", "hello", "world"},
				Env:  []string{"GODEBUG=1", "GOTRACEBACK=2", "KEY1=value1", "KEY2=value2"},
			},
			expected: []string{"KEY1=value1", "KEY2=value2"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			CleanCmdEnv(tc.cmd)
			assert.Equal(t, tc.expected, tc.cmd.Env)
		})
	}
}
