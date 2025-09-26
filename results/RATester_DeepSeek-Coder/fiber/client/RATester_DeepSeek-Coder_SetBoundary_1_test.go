package client

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetBoundary_1(t *testing.T) {
	type args struct {
		b string
	}
	tests := []struct {
		name string
		r    *Request
		args args
		want *Request
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

			if got := tt.r.SetBoundary(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.SetBoundary() = %v, want %v", got, tt.want)
			}
		})
	}
}
