package client

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestgetRetryConfig_1(t *testing.T) {
	type fields struct {
		client *Client
		req    *Request
		ctx    context.Context
	}
	tests := []struct {
		name   string
		fields fields
		want   *RetryConfig
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

			c := &core{
				client: tt.fields.client,
				req:    tt.fields.req,
				ctx:    tt.fields.ctx,
			}
			if got := c.getRetryConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("core.getRetryConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
