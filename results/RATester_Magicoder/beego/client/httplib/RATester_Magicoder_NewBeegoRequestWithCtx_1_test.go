package httplib

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestNewBeegoRequestWithCtx_1(t *testing.T) {
	type args struct {
		ctx    context.Context
		rawurl string
		method string
	}
	tests := []struct {
		name string
		args args
		want *BeegoHTTPRequest
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

			if got := NewBeegoRequestWithCtx(tt.args.ctx, tt.args.rawurl, tt.args.method); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBeegoRequestWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
