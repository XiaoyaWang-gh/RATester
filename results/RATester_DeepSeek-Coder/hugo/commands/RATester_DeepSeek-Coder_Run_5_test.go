package commands

import (
	"context"
	"fmt"
	"testing"

	"github.com/bep/simplecobra"
)

func TestRun_5(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		command *serverCommand
		wantErr bool
	}{
		{
			name: "Test case 1",
			command: &serverCommand{
				pprof: true,
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			command: &serverCommand{
				serverWatch: true,
			},
			wantErr: false,
		},
		{
			name: "Test case 3",
			command: &serverCommand{
				pprof:       true,
				serverWatch: true,
			},
			wantErr: false,
		},
		{
			name: "Test case 4",
			command: &serverCommand{
				serverWatch: true,
				pprof:       true,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ctx := context.Background()
			cd := &simplecobra.Commandeer{
				Command: tt.command,
			}
			args := []string{}
			err := tt.command.Run(ctx, cd, args)
			if (err != nil) != tt.wantErr {
				t.Errorf("serverCommand.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
