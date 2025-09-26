package client

import (
	"fmt"
	"testing"
)

func TestURL_1(t *testing.T) {
	type fields struct {
		url string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "TestURL",
			fields: fields{
				url: "http://example.com",
			},
			want: "http://example.com",
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
				url: tt.fields.url,
			}
			if got := r.URL(); got != tt.want {
				t.Errorf("Request.URL() = %v, want %v", got, tt.want)
			}
		})
	}
}
