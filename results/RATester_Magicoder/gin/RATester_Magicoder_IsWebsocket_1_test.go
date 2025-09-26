package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestIsWebsocket_1(t *testing.T) {
	tests := []struct {
		name string
		ctx  *Context
		want bool
	}{
		{
			name: "Test case 1",
			ctx: &Context{
				Request: &http.Request{
					Header: http.Header{
						"Connection": []string{"Upgrade"},
						"Upgrade":    []string{"websocket"},
					},
				},
			},
			want: true,
		},
		{
			name: "Test case 2",
			ctx: &Context{
				Request: &http.Request{
					Header: http.Header{
						"Connection": []string{"Upgrade"},
						"Upgrade":    []string{"not websocket"},
					},
				},
			},
			want: false,
		},
		{
			name: "Test case 3",
			ctx: &Context{
				Request: &http.Request{
					Header: http.Header{
						"Connection": []string{"not Upgrade"},
						"Upgrade":    []string{"websocket"},
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

			if got := tt.ctx.IsWebsocket(); got != tt.want {
				t.Errorf("IsWebsocket() = %v, want %v", got, tt.want)
			}
		})
	}
}
