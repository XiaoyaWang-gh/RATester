package client

import (
	"fmt"
	"net"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestNewStaticFilePlugin_1(t *testing.T) {
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
			name: "test",
			args: args{
				options: &v1.StaticFilePluginOptions{
					LocalPath: "test",
				},
			},
			want: &StaticFilePlugin{
				opts: &v1.StaticFilePluginOptions{
					LocalPath: "test",
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

			got, err := NewStaticFilePlugin(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewStaticFilePlugin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStaticFilePlugin() got = %v, want %v", got, tt.want)
			}
		})
	}
}
