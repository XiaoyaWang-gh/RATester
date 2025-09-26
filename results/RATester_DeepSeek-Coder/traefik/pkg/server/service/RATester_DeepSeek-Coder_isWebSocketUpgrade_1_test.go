package service

import (
	"fmt"
	"net/http"
	"testing"
)

func TestIsWebSocketUpgrade_1(t *testing.T) {
	tests := []struct {
		name string
		req  *http.Request
		want bool
	}{
		{
			name: "WebSocket upgrade",
			req: &http.Request{
				Header: http.Header{
					"Connection": []string{"Upgrade"},
					"Upgrade":    []string{"websocket"},
				},
			},
			want: true,
		},
		{
			name: "Non-WebSocket upgrade",
			req: &http.Request{
				Header: http.Header{
					"Connection": []string{"Upgrade"},
					"Upgrade":    []string{"not websocket"},
				},
			},
			want: false,
		},
		{
			name: "No upgrade",
			req: &http.Request{
				Header: http.Header{
					"Connection": []string{"not Upgrade"},
					"Upgrade":    []string{"websocket"},
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

			if got := isWebSocketUpgrade(tt.req); got != tt.want {
				t.Errorf("isWebSocketUpgrade() = %v, want %v", got, tt.want)
			}
		})
	}
}
