package metrics

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetRequestProtocol_1(t *testing.T) {
	tests := []struct {
		name string
		req  *http.Request
		want string
	}{
		{
			name: "Websocket Request",
			req: &http.Request{
				Header: http.Header{
					"Upgrade": []string{"websocket"},
				},
			},
			want: protoWebsocket,
		},
		{
			name: "SSE Request",
			req: &http.Request{
				Header: http.Header{
					"Accept": []string{"text/event-stream"},
				},
			},
			want: protoSSE,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getRequestProtocol(tt.req); got != tt.want {
				t.Errorf("getRequestProtocol() = %v, want %v", got, tt.want)
			}
		})
	}
}
