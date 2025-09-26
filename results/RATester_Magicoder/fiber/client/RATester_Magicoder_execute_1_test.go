package client

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func Testexecute_1(t *testing.T) {
	type args struct {
		ctx    context.Context
		client *Client
		req    *Request
	}
	tests := []struct {
		name    string
		c       *core
		args    args
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

			got, err := tt.c.execute(tt.args.ctx, tt.args.client, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("core.execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("core.execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
