package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRouter_3(t *testing.T) {
	type args struct {
		rootpath       string
		c              ControllerInterface
		mappingMethods []string
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

			if got := Router(tt.args.rootpath, tt.args.c, tt.args.mappingMethods...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Router() = %v, want %v", got, tt.want)
			}
		})
	}
}
