package net

import (
	"context"
	"fmt"
	"net"
	"testing"
)

func TestWithContext_1(t *testing.T) {
	type fields struct {
		Conn net.Conn
		Ctx  context.Context
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   context.Context
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

			c := &ContextConn{
				Conn: tt.fields.Conn,
				ctx:  tt.fields.Ctx,
			}
			c.WithContext(tt.args.ctx)
			if c.ctx != tt.want {
				t.Errorf("WithContext() = %v, want %v", c.ctx, tt.want)
			}
		})
	}
}
