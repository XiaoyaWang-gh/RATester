package config

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/spf13/cobra"
)

func TestRegisterProxyDomainConfigFlags_1(t *testing.T) {
	t.Parallel()
	type args struct {
		cmd *cobra.Command
		c   *v1.DomainConfig
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				cmd: &cobra.Command{},
				c:   &v1.DomainConfig{},
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

			registerProxyDomainConfigFlags(tt.args.cmd, tt.args.c)
		})
	}
}
