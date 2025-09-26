package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewNode_1(t *testing.T) {
	type args struct {
		t               kind
		pre             string
		p               *node
		sc              children
		originalPath    string
		methods         *routeMethods
		paramsCount     int
		paramChildren   *node
		anyChildren     *node
		notFoundHandler *routeMethod
	}
	tests := []struct {
		name string
		args args
		want *node
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

			if got := newNode(tt.args.t, tt.args.pre, tt.args.p, tt.args.sc, tt.args.originalPath, tt.args.methods, tt.args.paramsCount, tt.args.paramChildren, tt.args.anyChildren, tt.args.notFoundHandler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
