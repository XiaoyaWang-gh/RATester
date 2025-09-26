package metrics

import (
	"fmt"
	"net/http"
	"testing"
)

func TestIsSSERequest_1(t *testing.T) {
	tests := []struct {
		name string
		req  *http.Request
		want bool
	}{
		{
			name: "test case 1",
			req: &http.Request{
				Header: http.Header{
					"Accept": []string{"text/event-stream"},
				},
			},
			want: true,
		},
		{
			name: "test case 2",
			req: &http.Request{
				Header: http.Header{
					"Accept": []string{"text/event-stream", "text/html"},
				},
			},
			want: true,
		},
		{
			name: "test case 3",
			req: &http.Request{
				Header: http.Header{
					"Accept": []string{"text/html"},
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

			if got := isSSERequest(tt.req); got != tt.want {
				t.Errorf("isSSERequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
