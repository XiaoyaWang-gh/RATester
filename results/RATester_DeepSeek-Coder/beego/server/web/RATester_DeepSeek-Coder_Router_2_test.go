package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRouter_2(t *testing.T) {
	type args struct {
		rootpath       string
		c              ControllerInterface
		mappingMethods []string
	}
	tests := []struct {
		name string
		n    *Namespace
		args args
		want *Namespace
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

			if got := tt.n.Router(tt.args.rootpath, tt.args.c, tt.args.mappingMethods...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Namespace.Router() = %v, want %v", got, tt.want)
			}
		})
	}
}
