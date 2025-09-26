package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInsertFilterChain_2(t *testing.T) {
	type args struct {
		pattern     string
		filterChain FilterChain
		opts        []FilterOpt
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

			if got := InsertFilterChain(tt.args.pattern, tt.args.filterChain, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertFilterChain() = %v, want %v", got, tt.want)
			}
		})
	}
}
