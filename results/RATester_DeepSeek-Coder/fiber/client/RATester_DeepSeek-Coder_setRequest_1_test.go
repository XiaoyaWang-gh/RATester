package client

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetRequest_1(t *testing.T) {
	type args struct {
		req *Request
	}
	tests := []struct {
		name string
		r    *Response
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

			r := &Response{
				client:  tt.r.client,
				request: tt.r.request,
			}
			r.setRequest(tt.args.req)
			if !reflect.DeepEqual(r.request, tt.want) {
				t.Errorf("Response.setRequest() = %v, want %v", r.request, tt.want)
			}
		})
	}
}
