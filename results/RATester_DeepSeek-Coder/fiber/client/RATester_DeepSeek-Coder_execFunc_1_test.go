package client

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestExecFunc_1(t *testing.T) {
	type fields struct {
		client *Client
		req    *Request
		ctx    context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Response
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
			got, err := c.execFunc()
			if (err != nil) != tt.wantErr {
				t.Errorf("core.execFunc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("core.execFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
