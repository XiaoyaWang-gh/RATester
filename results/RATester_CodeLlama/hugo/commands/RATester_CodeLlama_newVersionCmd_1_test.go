package commands

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/bep/simplecobra"
	"github.com/gohugoio/hugo/common/hugo"
	"github.com/spf13/cobra"
)

func TestNewVersionCmd_1(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		args args
		want simplecobra.Commander
	}{
		{
			name: "test_newVersionCmd",
			args: args{},
			want: &simpleCommand{
				name: "version",
				run: func(ctx context.Context, cd *simplecobra.Commandeer, r *rootCommand, args []string) error {
					r.Println(hugo.BuildVersionString())
					return nil
				},
				short: "Print Hugo version and environment info",
				long:  "Print Hugo version and environment info. This is useful in Hugo bug reports.",
				withc: func(cmd *cobra.Command, r *rootCommand) {
					cmd.ValidArgsFunction = cobra.NoFileCompletions
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := newVersionCmd(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newVersionCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}
