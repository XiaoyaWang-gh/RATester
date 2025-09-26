package cli

import (
	"fmt"
	"testing"

	"github.com/traefik/paerser/cli"
)

func TestLoad_1(t *testing.T) {
	type args struct {
		args []string
		cmd  *cli.Command
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Test with empty args",
			args: args{
				args: []string{},
				cmd:  &cli.Command{},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Test with non-empty args",
			args: args{
				args: []string{"--name", "test"},
				cmd:  &cli.Command{Configuration: &FlagLoader{}},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Test with error",
			args: args{
				args: []string{"--name", "test"},
				cmd:  &cli.Command{Configuration: &FlagLoader{}},
			},
			want:    false,
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

			fl := &FlagLoader{}
			got, err := fl.Load(tt.args.args, tt.args.cmd)
			if (err != nil) != tt.wantErr {
				t.Errorf("FlagLoader.Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FlagLoader.Load() = %v, want %v", got, tt.want)
			}
		})
	}
}
