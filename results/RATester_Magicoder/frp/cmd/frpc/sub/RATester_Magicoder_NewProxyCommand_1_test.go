package sub

import (
	"fmt"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/spf13/cobra"
)

func TestNewProxyCommand_1(t *testing.T) {
	type args struct {
		name      string
		c         v1.ProxyConfigurer
		clientCfg *v1.ClientCommonConfig
	}
	tests := []struct {
		name string
		args args
		want *cobra.Command
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

			if got := NewProxyCommand(tt.args.name, tt.args.c, tt.args.clientCfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProxyCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
