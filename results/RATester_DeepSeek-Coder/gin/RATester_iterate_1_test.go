package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIterate_1(t *testing.T) {
	type args struct {
		path   string
		method string
		routes RoutesInfo
		root   *node
	}
	tests := []struct {
		name string
		args args
		want RoutesInfo
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

			if got := iterate(tt.args.path, tt.args.method, tt.args.routes, tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("iterate() = %v, want %v", got, tt.want)
			}
		})
	}
}
