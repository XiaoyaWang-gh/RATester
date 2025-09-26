package api

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/runtime"
)

func TestNewRouterRepresentation_1(t *testing.T) {
	type args struct {
		name string
		rt   *runtime.RouterInfo
	}
	tests := []struct {
		name string
		args args
		want routerRepresentation
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

			if got := newRouterRepresentation(tt.args.name, tt.args.rt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newRouterRepresentation() = %v, want %v", got, tt.want)
			}
		})
	}
}
