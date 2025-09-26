package client

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestResetPathParams_1(t *testing.T) {
	type fields struct {
		ctx  context.Context
		path *PathParam
	}
	tests := []struct {
		name   string
		fields fields
		want   *Request
	}{
		{
			name: "TestResetPathParams",
			fields: fields{
				ctx: context.Background(),
				path: &PathParam{
					"key": "value",
				},
			},
			want: &Request{
				ctx: context.Background(),
				path: &PathParam{
					"key": "value",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			r := &Request{
				ctx:  tt.fields.ctx,
				path: tt.fields.path,
			}
			if got := r.ResetPathParams(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResetPathParams() = %v, want %v", got, tt.want)
			}
		})
	}
}
