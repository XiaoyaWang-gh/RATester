package client

import (
	"fmt"
	"testing"
)

func TestMethod_2(t *testing.T) {
	type fields struct {
		method string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "TestMethod_Get",
			fields: fields{
				method: "GET",
			},
			want: "GET",
		},
		{
			name: "TestMethod_Post",
			fields: fields{
				method: "POST",
			},
			want: "POST",
		},
		{
			name: "TestMethod_Put",
			fields: fields{
				method: "PUT",
			},
			want: "PUT",
		},
		{
			name: "TestMethod_Delete",
			fields: fields{
				method: "DELETE",
			},
			want: "DELETE",
		},
		{
			name: "TestMethod_Patch",
			fields: fields{
				method: "PATCH",
			},
			want: "PATCH",
		},
		{
			name: "TestMethod_Options",
			fields: fields{
				method: "OPTIONS",
			},
			want: "OPTIONS",
		},
		{
			name: "TestMethod_Head",
			fields: fields{
				method: "HEAD",
			},
			want: "HEAD",
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
				method: tt.fields.method,
			}
			if got := r.Method(); got != tt.want {
				t.Errorf("Request.Method() = %v, want %v", got, tt.want)
			}
		})
	}
}
