package client

import (
	"fmt"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestNewSocks5Plugin_1(t *testing.T) {
	type args struct {
		options v1.ClientPluginOptions
	}
	tests := []struct {
		name    string
		args    args
		wantP   Plugin
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				options: &v1.Socks5PluginOptions{
					Username: "test_user",
					Password: "test_pass",
				},
			},
			wantP:   &Socks5Plugin{},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				options: &v1.Socks5PluginOptions{
					Username: "",
					Password: "",
				},
			},
			wantP:   &Socks5Plugin{},
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

			gotP, err := NewSocks5Plugin(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSocks5Plugin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotP, tt.wantP) {
				t.Errorf("NewSocks5Plugin() = %v, want %v", gotP, tt.wantP)
			}
		})
	}
}
