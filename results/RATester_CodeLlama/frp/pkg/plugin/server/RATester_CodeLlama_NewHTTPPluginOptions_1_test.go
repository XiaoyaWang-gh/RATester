package plugin

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestNewHTTPPluginOptions_1(t *testing.T) {
	type args struct {
		options v1.HTTPPluginOptions
	}
	tests := []struct {
		name string
		args args
		want Plugin
	}{
		{
			name: "test case 1",
			args: args{
				options: v1.HTTPPluginOptions{
					Name:      "test",
					Addr:      "127.0.0.1",
					Path:      "/test",
					Ops:       []string{"test"},
					TLSVerify: true,
				},
			},
			want: &httpPlugin{
				options: v1.HTTPPluginOptions{
					Name:      "test",
					Addr:      "127.0.0.1",
					Path:      "/test",
					Ops:       []string{"test"},
					TLSVerify: true,
				},
				url:    "https://127.0.0.1/test",
				client: &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: false}}},
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

			if got := NewHTTPPluginOptions(tt.args.options); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHTTPPluginOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}
