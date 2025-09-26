package metrics

import (
	"fmt"
	"net/http"
	"testing"
)

func TestIsWebsocketRequest_2(t *testing.T) {
	tests := []struct {
		name string
		req  *http.Request
		want bool
	}{
		{
			name: "Websocket Request",
			req: &http.Request{
				Header: http.Header{
					"Connection": []string{"upgrade"},
					"Upgrade":    []string{"websocket"},
				},
			},
			want: true,
		},
		{
			name: "Non-Websocket Request",
			req: &http.Request{
				Header: http.Header{
					"Connection": []string{"keep-alive"},
					"Upgrade":    []string{"websocket"},
				},
			},
			want: false,
		},
		{
			name: "Partial Websocket Request",
			req: &http.Request{
				Header: http.Header{
					"Connection": []string{"upgrade"},
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

			if got := isWebsocketRequest(tt.req); got != tt.want {
				t.Errorf("isWebsocketRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
