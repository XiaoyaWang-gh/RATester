package middleware

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/runtime"
)

func TestNewBuilder_2(t *testing.T) {
	type args struct {
		configs        map[string]*runtime.MiddlewareInfo
		serviceBuilder serviceBuilder
		pluginBuilder  PluginsBuilder
	}
	tests := []struct {
		name string
		args args
		want *Builder
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

			if got := NewBuilder(tt.args.configs, tt.args.serviceBuilder, tt.args.pluginBuilder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}
