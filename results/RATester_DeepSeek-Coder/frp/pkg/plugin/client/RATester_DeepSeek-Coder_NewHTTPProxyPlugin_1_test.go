package client

import (
	"fmt"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestNewHTTPProxyPlugin_1(t *testing.T) {
	type args struct {
		options v1.ClientPluginOptions
	}
	tests := []struct {
		name    string
		args    args
		want    Plugin
		wantErr bool
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

			got, err := NewHTTPProxyPlugin(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHTTPProxyPlugin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHTTPProxyPlugin() = %v, want %v", got, tt.want)
			}
		})
	}
}
