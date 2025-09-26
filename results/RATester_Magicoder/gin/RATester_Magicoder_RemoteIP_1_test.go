package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRemoteIP_1(t *testing.T) {
	tests := []struct {
		name string
		ctx  *Context
		want string
	}{
		{
			name: "Test with valid remote address",
			ctx: &Context{
				Request: &http.Request{
					RemoteAddr: "192.0.2.1:1234",
				},
			},
			want: "192.0.2.1",
		},
		{
			name: "Test with invalid remote address",
			ctx: &Context{
				Request: &http.Request{
					RemoteAddr: "invalid",
				},
			},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.ctx.RemoteIP(); got != tt.want {
				t.Errorf("RemoteIP() = %v, want %v", got, tt.want)
			}
		})
	}
}
