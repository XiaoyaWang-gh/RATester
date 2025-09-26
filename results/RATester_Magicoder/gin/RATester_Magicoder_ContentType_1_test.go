package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestContentType_1(t *testing.T) {
	tests := []struct {
		name string
		ctx  *Context
		want string
	}{
		{
			name: "Test case 1",
			ctx: &Context{
				Request: &http.Request{
					Header: http.Header{
						"Content-Type": []string{"application/json"},
					},
				},
			},
			want: "application/json",
		},
		{
			name: "Test case 2",
			ctx: &Context{
				Request: &http.Request{
					Header: http.Header{
						"Content-Type": []string{"text/plain"},
					},
				},
			},
			want: "text/plain",
		},
		{
			name: "Test case 3",
			ctx: &Context{
				Request: &http.Request{
					Header: http.Header{
						"Content-Type": []string{"application/xml"},
					},
				},
			},
			want: "application/xml",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.ctx.ContentType(); got != tt.want {
				t.Errorf("ContentType() = %v, want %v", got, tt.want)
			}
		})
	}
}
