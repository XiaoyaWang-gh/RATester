package client

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetMethod_1(t *testing.T) {
	type args struct {
		method string
	}
	tests := []struct {
		name string
		r    *Request
		args args
		want *Request
	}{
		{
			name: "TestSetMethod_GET",
			r:    &Request{},
			args: args{method: "GET"},
			want: &Request{method: "GET"},
		},
		{
			name: "TestSetMethod_POST",
			r:    &Request{},
			args: args{method: "POST"},
			want: &Request{method: "POST"},
		},
		{
			name: "TestSetMethod_PUT",
			r:    &Request{},
			args: args{method: "PUT"},
			want: &Request{method: "PUT"},
		},
		{
			name: "TestSetMethod_DELETE",
			r:    &Request{},
			args: args{method: "DELETE"},
			want: &Request{method: "DELETE"},
		},
		{
			name: "TestSetMethod_PATCH",
			r:    &Request{},
			args: args{method: "PATCH"},
			want: &Request{method: "PATCH"},
		},
		{
			name: "TestSetMethod_OPTIONS",
			r:    &Request{},
			args: args{method: "OPTIONS"},
			want: &Request{method: "OPTIONS"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.r.SetMethod(tt.args.method); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.SetMethod() = %v, want %v", got, tt.want)
			}
		})
	}
}
