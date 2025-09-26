package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUnregisterFixedRoute_2(t *testing.T) {
	type args struct {
		fixedRoute string
		method     string
	}
	tests := []struct {
		name string
		args args
		want *HttpServer
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

			if got := UnregisterFixedRoute(tt.args.fixedRoute, tt.args.method); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnregisterFixedRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}
