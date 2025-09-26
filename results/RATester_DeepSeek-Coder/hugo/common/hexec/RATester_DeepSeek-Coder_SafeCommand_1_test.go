package hexec

import (
	"fmt"
	"testing"
)

func TestSafeCommand_1(t *testing.T) {
	testCases := []struct {
		name    string
		command string
		args    []string
		wantErr bool
	}{
		{
			name:    "valid command",
			command: "ls",
			args:    []string{"-l"},
			wantErr: false,
		},
		{
			name:    "invalid command",
			command: "invalid_command",
			args:    []string{"-l"},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, err := SafeCommand(tc.command, tc.args...)
			if (err != nil) != tc.wantErr {
				t.Errorf("SafeCommand() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
