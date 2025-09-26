package client

import (
	"fmt"
	"net"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestNewUnixDomainSocketPlugin_1(t *testing.T) {
	type args struct {
		options v1.ClientPluginOptions
	}
	tests := []struct {
		name    string
		args    args
		want    Plugin
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				options: &v1.UnixDomainSocketPluginOptions{
					UnixPath: "/tmp/frp.sock",
				},
			},
			want: &UnixDomainSocketPlugin{
				UnixAddr: &net.UnixAddr{
					Name: "/tmp/frp.sock",
					Net:  "unix",
				},
			},
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

			got, err := NewUnixDomainSocketPlugin(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUnixDomainSocketPlugin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUnixDomainSocketPlugin() got = %v, want %v", got, tt.want)
			}
		})
	}
}
