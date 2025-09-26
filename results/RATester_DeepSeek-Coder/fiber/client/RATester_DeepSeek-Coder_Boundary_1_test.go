package client

import (
	"fmt"
	"testing"
)

func TestBoundary_1(t *testing.T) {
	type fields struct {
		boundary string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "TestBoundary",
			fields: fields{
				boundary: "testBoundary",
			},
			want: "testBoundary",
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
				boundary: tt.fields.boundary,
			}
			if got := r.Boundary(); got != tt.want {
				t.Errorf("Request.Boundary() = %v, want %v", got, tt.want)
			}
		})
	}
}
