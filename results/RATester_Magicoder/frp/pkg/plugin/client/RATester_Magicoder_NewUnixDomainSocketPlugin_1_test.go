package client

import (
	"fmt"
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
		wantP   Plugin
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

			gotP, err := NewUnixDomainSocketPlugin(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUnixDomainSocketPlugin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotP, tt.wantP) {
				t.Errorf("NewUnixDomainSocketPlugin() = %v, want %v", gotP, tt.wantP)
			}
		})
	}
}
