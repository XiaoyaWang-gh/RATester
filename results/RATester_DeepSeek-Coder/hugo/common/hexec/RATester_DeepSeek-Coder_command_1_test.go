package hexec

import (
	"fmt"
	"testing"
)

func TestCommand_1(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name    string
		command *commandeer
		args    []any
		wantErr bool
	}

	tests := []testCase{
		{
			name: "valid command",
			command: &commandeer{
				name: "echo",
			},
			args:    []any{"hello", "world"},
			wantErr: false,
		},
		{
			name: "invalid command",
			command: &commandeer{
				name: "invalid_command",
			},
			args:    []any{"hello", "world"},
			wantErr: true,
		},
		{
			name:    "nil commandeer",
			args:    []any{"hello", "world"},
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

			_, err := tt.command.command(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("command() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
