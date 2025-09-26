package client

import (
	"fmt"
	"testing"
)

func TestUserAgent_1(t *testing.T) {
	type fields struct {
		userAgent string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "TestUserAgent",
			fields: fields{
				userAgent: "TestUserAgent",
			},
			want: "TestUserAgent",
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
				userAgent: tt.fields.userAgent,
			}
			if got := r.UserAgent(); got != tt.want {
				t.Errorf("UserAgent() = %v, want %v", got, tt.want)
			}
		})
	}
}
