package sub

import (
	"fmt"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/spf13/cobra"
)

func TestNewAdminCommand_1(t *testing.T) {
	type args struct {
		name    string
		short   string
		handler func(*v1.ClientCommonConfig) error
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

			got := NewAdminCommand(tt.args.name, tt.args.short, tt.args.handler)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAdminCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
