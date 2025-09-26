package client

import (
	"fmt"
	"net"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestNewHTTP2HTTPPlugin_1(t *testing.T) {
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
				options: &v1.HTTP2HTTPPluginOptions{
					LocalAddr: "127.0.0.1:8080",
				},
			},
			want: &HTTP2HTTPPlugin{
				opts: &v1.HTTP2HTTPPluginOptions{
					LocalAddr: "127.0.0.1:8080",
				},
				l: &Listener{
					conns: make(chan net.Conn, 10),
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

			got, err := NewHTTP2HTTPPlugin(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHTTP2HTTPPlugin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHTTP2HTTPPlugin() got = %v, want %v", got, tt.want)
			}
		})
	}
}
