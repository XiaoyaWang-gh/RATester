package plugin

import (
	"fmt"
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
		// TODO: Add test cases.
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
