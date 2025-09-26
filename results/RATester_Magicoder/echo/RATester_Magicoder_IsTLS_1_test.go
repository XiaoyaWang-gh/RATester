package echo

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"testing"
)

func TestIsTLS_1(t *testing.T) {
	type args struct {
		c *context
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				c: &context{
					request: &http.Request{
						TLS: &tls.ConnectionState{},
					},
				},
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				c: &context{
					request: &http.Request{
						TLS: nil,
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.args.c.IsTLS(); got != tt.want {
				t.Errorf("IsTLS() = %v, want %v", got, tt.want)
			}
		})
	}
}
