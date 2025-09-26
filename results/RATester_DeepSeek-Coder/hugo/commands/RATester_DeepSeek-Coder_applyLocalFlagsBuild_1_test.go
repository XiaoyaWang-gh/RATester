package commands

import (
	"fmt"
	"testing"

	"github.com/spf13/cobra"
)

func TestApplyLocalFlagsBuild_1(t *testing.T) {
	type args struct {
		cmd *cobra.Command
		r   *rootCommand
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			applyLocalFlagsBuild(tt.args.cmd, tt.args.r)
		})
	}
}
