package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRESTRouter_2(t *testing.T) {
	type args struct {
		rootpath string
		c        ControllerInterface
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

			if got := RESTRouter(tt.args.rootpath, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RESTRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}
