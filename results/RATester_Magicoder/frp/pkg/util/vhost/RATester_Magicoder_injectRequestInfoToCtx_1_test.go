package vhost

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestInjectRequestInfoToCtx_1(t *testing.T) {
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name string
		args args
		want *http.Request
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

			rp := &HTTPReverseProxy{}
			if got := rp.injectRequestInfoToCtx(tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("injectRequestInfoToCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
