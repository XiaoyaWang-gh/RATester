package hexec

import (
	"fmt"
	"testing"
)

func TestSafeCommand_1(t *testing.T) {
	tests := []struct {
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, err := SafeCommand(tt.command, tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("SafeCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
