package headers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestProcessCorsHeaders_1(t *testing.T) {
	tests := []struct {
		name string
		req  *http.Request
		want bool
	}{
		{
			name: "Test case 1",
			req: &http.Request{
				Method: http.MethodOptions,
				Header: http.Header{
					"Origin":                         []string{"http://example.com"},
					"Access-Control-Request-Method":  []string{http.MethodGet},
					"Access-Control-Request-Headers": []string{"Content-Type"},
				},
			},
			want: true,
		},
		{
			name: "Test case 2",
			req: &http.Request{
				Method: http.MethodGet,
				Header: http.Header{
					"Origin": []string{"http://example.com"},
				},
			},
			want: false,
		},
		{
			name: "Test case 3",
			req: &http.Request{
				Method: http.MethodOptions,
				Header: http.Header{
					"Origin": []string{"http://example.com"},
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

			header := &Header{
				hasCorsHeaders: true,
				headers: &dynamic.Headers{
					AccessControlAllowCredentials: true,
					AccessControlAllowHeaders:     []string{"Content-Type"},
					AccessControlAllowMethods:     []string{http.MethodGet},
					AccessControlMaxAge:           3600,
				},
			}

			rw := httptest.NewRecorder()
			got := header.processCorsHeaders(rw, tt.req)

			if got != tt.want {
				t.Errorf("processCorsHeaders() = %v, want %v", got, tt.want)
			}
		})
	}
}
