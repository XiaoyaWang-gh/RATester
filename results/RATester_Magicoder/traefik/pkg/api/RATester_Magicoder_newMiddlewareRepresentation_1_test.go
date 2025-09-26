package api

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/runtime"
)

func TestnewMiddlewareRepresentation_1(t *testing.T) {
	type args struct {
		name string
		mi   *runtime.MiddlewareInfo
	}
	tests := []struct {
		name string
		args args
		want middlewareRepresentation
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

			if got := newMiddlewareRepresentation(tt.args.name, tt.args.mi); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMiddlewareRepresentation() = %v, want %v", got, tt.want)
			}
		})
	}
}
