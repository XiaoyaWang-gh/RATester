package server

import (
	"fmt"
	"net"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/static"
)

func TestNewListenConfig_1(t *testing.T) {
	type args struct {
		configuration *static.EntryPoint
	}
	tests := []struct {
		name string
		args args
		want net.ListenConfig
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

			if got := newListenConfig(tt.args.configuration); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newListenConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
