package client

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestRegister_3(t *testing.T) {
	type args struct {
		name string
		fn   CreatorFn
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestRegister_0",
			args: args{
				name: "plugin_0",
				fn:   func(options v1.ClientPluginOptions) (Plugin, error) { return nil, nil },
			},
			want: "plugin [plugin_0] is already registered",
		},
		{
			name: "TestRegister_1",
			args: args{
				name: "plugin_1",
				fn:   func(options v1.ClientPluginOptions) (Plugin, error) { return nil, nil },
			},
			want: "plugin [plugin_1] is already registered",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			defer func() {
				if r := recover(); r == nil {
					t.Errorf("The code did not panic")
				} else if r != tt.want {
					t.Errorf("Register() = %v, want %v", r, tt.want)
				}
			}()
			Register(tt.args.name, tt.args.fn)
		})
	}
}
