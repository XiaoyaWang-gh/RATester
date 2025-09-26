package client

import (
	"fmt"
	"testing"
)

func TestReferer_1(t *testing.T) {
	type fields struct {
		referer string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "TestReferer",
			fields: fields{
				referer: "http://example.com",
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
				referer: tt.fields.referer,
			}
			if got := r.Referer(); got != tt.want {
				t.Errorf("Request.Referer() = %v, want %v", got, tt.want)
			}
		})
	}
}
