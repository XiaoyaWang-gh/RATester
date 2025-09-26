package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestnewFilterRouter_1(t *testing.T) {
	type args struct {
		pattern string
		filter  FilterFunc
		opts    []FilterOpt
	}
	tests := []struct {
		name string
		args args
		want *FilterRouter
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

			if got := newFilterRouter(tt.args.pattern, tt.args.filter, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newFilterRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}
