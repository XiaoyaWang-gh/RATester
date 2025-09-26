package client

import (
	"context"
	"fmt"
	"testing"
)

func TestPreHooks_1(t *testing.T) {
	type fields struct {
		client *Client
		req    *Request
		ctx    context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
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
			if err := c.preHooks(); (err != nil) != tt.wantErr {
				t.Errorf("core.preHooks() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
