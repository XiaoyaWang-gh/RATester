package client

import (
	"fmt"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestNewHTTPS2HTTPSPlugin_1(t *testing.T) {
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
				options: &v1.HTTPS2HTTPSPluginOptions{
					LocalAddr: "127.0.0.1:8080",
				},
			},
			want: &HTTPS2HTTPSPlugin{
				opts: &v1.HTTPS2HTTPSPluginOptions{
					LocalAddr: "127.0.0.1:8080",
				},
				l: NewProxyListener(),
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

			got, err := NewHTTPS2HTTPSPlugin(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHTTPS2HTTPSPlugin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHTTPS2HTTPSPlugin() got = %v, want %v", got, tt.want)
			}
		})
	}
}
