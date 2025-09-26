package client

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestContext_2(t *testing.T) {
	type fields struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		want   context.Context
	}{
		{
			name: "Test Case 1",
			fields: fields{
				ctx: context.Background(),
			},
			want: context.Background(),
		},
		{
			name: "Test Case 2",
			fields: fields{
				ctx: context.TODO(),
			},
			want: context.TODO(),
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
				ctx: tt.fields.ctx,
			}
			if got := r.Context(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.Context() = %v, want %v", got, tt.want)
			}
		})
	}
}
