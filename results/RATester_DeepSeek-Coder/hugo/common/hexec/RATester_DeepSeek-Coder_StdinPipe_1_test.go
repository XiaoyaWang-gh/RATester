package hexec

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestStdinPipe_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		command string
		wantErr bool
	}{
		{
			name:    "success",
			command: "echo",
			wantErr: false,
		},
		{
			name:    "failure",
			command: "non-existent-command",
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			c := &cmdWrapper{
				name: tc.name,
				c:    exec.Command(tc.command),
			}

			_, err := c.StdinPipe()
			if (err != nil) != tc.wantErr {
				t.Errorf("StdinPipe() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
