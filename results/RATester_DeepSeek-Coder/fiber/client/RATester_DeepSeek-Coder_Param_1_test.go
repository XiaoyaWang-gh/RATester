package client

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestParam_1(t *testing.T) {
	type fields struct {
		ctx    context.Context
		params *QueryParam
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes []string
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

			r := &Request{
				ctx:    tt.fields.ctx,
				params: tt.fields.params,
			}
			if gotRes := r.Param(tt.args.key); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Request.Param() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
