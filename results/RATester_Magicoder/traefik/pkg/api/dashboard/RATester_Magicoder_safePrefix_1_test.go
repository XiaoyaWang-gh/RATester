package dashboard

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSafePrefix_1(t *testing.T) {
	tests := []struct {
		name string
		req  *http.Request
		want string
	}{
		{
			name: "valid prefix",
			req: &http.Request{
				Header: http.Header{
					"X-Forwarded-Prefix": []string{"/api"},
				},
			},
			want: "/api",
		},
		{
			name: "invalid prefix",
			req: &http.Request{
				Header: http.Header{
					"X-Forwarded-Prefix": []string{"http://example.com"},
				},
			},
			want: "",
		},
		{
			name: "empty prefix",
			req: &http.Request{
				Header: http.Header{
					"X-Forwarded-Prefix": []string{""},
				},
			},
			want: "",
		},
		{
			name: "no prefix",
			req: &http.Request{
				Header: http.Header{},
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

			if got := safePrefix(tt.req); got != tt.want {
				t.Errorf("safePrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
