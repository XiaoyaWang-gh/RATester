package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInsertFilter_3(t *testing.T) {
	type args struct {
		pattern string
		pos     int
		filter  FilterFunc
		opts    []FilterOpt
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

			if got := InsertFilter(tt.args.pattern, tt.args.pos, tt.args.filter, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}
