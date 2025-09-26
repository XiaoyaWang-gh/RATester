package config

import (
	"fmt"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/spf13/cobra"
)

func TestRegisterProxyBaseConfigFlags_1(t *testing.T) {
	type args struct {
		cmd  *cobra.Command
		c    *v1.ProxyBaseConfig
		opts []RegisterFlagOption
	}
	tests := []struct {
		name string
		args args
		want *v1.ProxyBaseConfig
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

			registerProxyBaseConfigFlags(tt.args.cmd, tt.args.c, tt.args.opts...)
			if !reflect.DeepEqual(tt.args.c, tt.want) {
				t.Errorf("registerProxyBaseConfigFlags() = %v, want %v", tt.args.c, tt.want)
			}
		})
	}
}
