package config

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/spf13/cobra"
)

func TestRegisterServerConfigFlags_1(t *testing.T) {
	type args struct {
		cmd  *cobra.Command
		c    *v1.ServerConfig
		opts []RegisterFlagOption
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

			RegisterServerConfigFlags(tt.args.cmd, tt.args.c, tt.args.opts...)
		})
	}
}
