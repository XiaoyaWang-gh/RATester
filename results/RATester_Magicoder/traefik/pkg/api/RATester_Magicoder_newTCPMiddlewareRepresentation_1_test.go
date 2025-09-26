package api

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/runtime"
)

func TestnewTCPMiddlewareRepresentation_1(t *testing.T) {
	type args struct {
		name string
		mi   *runtime.TCPMiddlewareInfo
	}
	tests := []struct {
		name string
		args args
		want tcpMiddlewareRepresentation
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

			if got := newTCPMiddlewareRepresentation(tt.args.name, tt.args.mi); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTCPMiddlewareRepresentation() = %v, want %v", got, tt.want)
			}
		})
	}
}
