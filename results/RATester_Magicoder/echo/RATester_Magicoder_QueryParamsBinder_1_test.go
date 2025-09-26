package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestQueryParamsBinder_1(t *testing.T) {
	type args struct {
		c Context
	}
	tests := []struct {
		name string
		args args
		want *ValueBinder
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

			if got := QueryParamsBinder(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryParamsBinder() = %v, want %v", got, tt.want)
			}
		})
	}
}
