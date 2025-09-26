package forwardedheaders

import (
	"fmt"
	"net/http"
	"testing"
)

func TestIsWebsocketRequest_1(t *testing.T) {
	tests := []struct {
		name string
		req  *http.Request
		want bool
	}{
		{
			name: "Websocket request",
			req: &http.Request{
				Header: http.Header{
					"Connection": []string{"upgrade"},
					"Upgrade":    []string{"websocket"},
				},
			},
			want: true,
		},
		{
			name: "Non-websocket request",
			req: &http.Request{
				Header: http.Header{
					"Connection": []string{"close"},
					"Upgrade":    []string{"websocket"},
				},
			},
			want: false,
		},
		{
			name: "Partial websocket request",
			req: &http.Request{
				Header: http.Header{
					"Connection": []string{"upgrade, close"},
					"Upgrade":    []string{"websocket"},
				},
			},
			want: true,
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
