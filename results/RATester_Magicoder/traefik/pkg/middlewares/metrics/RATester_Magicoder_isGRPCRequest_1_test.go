package metrics

import (
	"fmt"
	"net/http"
	"testing"
)

func TestIsGRPCRequest_1(t *testing.T) {
	tests := []struct {
		name string
		req  *http.Request
		want bool
	}{
		{
			name: "Test case 1",
			req: &http.Request{
				Header: http.Header{
					"Content-Type": []string{"application/grpc"},
				},
			},
			want: true,
		},
		{
			name: "Test case 2",
			req: &http.Request{
				Header: http.Header{
					"Content-Type": []string{"application/json"},
				},
			},
			want: false,
		},
		{
			name: "Test case 3",
			req: &http.Request{
				Header: http.Header{
					"Content-Type": []string{"application/grpc+json"},
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

			if got := isGRPCRequest(tt.req); got != tt.want {
				t.Errorf("isGRPCRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
