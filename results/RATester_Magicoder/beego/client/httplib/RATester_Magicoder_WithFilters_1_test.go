package httplib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWithFilters_1(t *testing.T) {
	type args struct {
		fcs []FilterChain
	}
	tests := []struct {
		name string
		args args
		want BeegoHTTPRequestOption
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

			if got := WithFilters(tt.args.fcs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithFilters() = %v, want %v", got, tt.want)
			}
		})
	}
}
