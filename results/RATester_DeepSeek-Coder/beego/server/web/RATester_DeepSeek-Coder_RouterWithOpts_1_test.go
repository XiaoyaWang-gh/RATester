package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRouterWithOpts_1(t *testing.T) {
	type args struct {
		rootpath string
		c        ControllerInterface
		opts     []ControllerOption
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

			if got := RouterWithOpts(tt.args.rootpath, tt.args.c, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RouterWithOpts() = %v, want %v", got, tt.want)
			}
		})
	}
}
