package client

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetClient_1(t *testing.T) {
	type args struct {
		c *Client
	}
	tests := []struct {
		name string
		r    *Response
		args args
		want *Client
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
				client: tt.r.client,
			}
			r.setClient(tt.args.c)
			if !reflect.DeepEqual(r.client, tt.want) {
				t.Errorf("Response.setClient() = %v, want %v", r.client, tt.want)
			}
		})
	}
}
