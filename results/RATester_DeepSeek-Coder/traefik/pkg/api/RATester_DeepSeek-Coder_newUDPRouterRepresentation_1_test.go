package api

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/runtime"
)

func TestNewUDPRouterRepresentation_1(t *testing.T) {
	type args struct {
		name string
		rt   *runtime.UDPRouterInfo
	}
	tests := []struct {
		name string
		args args
		want udpRouterRepresentation
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

			if got := newUDPRouterRepresentation(tt.args.name, tt.args.rt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newUDPRouterRepresentation() = %v, want %v", got, tt.want)
			}
		})
	}
}
