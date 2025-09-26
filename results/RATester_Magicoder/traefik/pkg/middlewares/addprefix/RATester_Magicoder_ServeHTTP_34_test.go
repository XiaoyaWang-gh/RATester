package addprefix

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestServeHTTP_34(t *testing.T) {
	tests := []struct {
		name         string
		req          *http.Request
		expectedPath string
		expectedRaw  string
		expectedURI  string
	}{
		{
			name: "Test with empty path",
			req: &http.Request{
				URL: &url.URL{
					Path: "",
				},
			},
			expectedPath: "/prefix/",
			expectedRaw:  "/prefix/",
			expectedURI:  "/prefix/",
		},
		{
			name: "Test with non-empty path",
			req: &http.Request{
				URL: &url.URL{
					Path: "/test",
				},
			},
			expectedPath: "/prefix/test",
			expectedRaw:  "/prefix/test",
			expectedURI:  "/prefix/test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			a := &addPrefix{
				prefix: "/prefix",
			}

			a.ServeHTTP(nil, tt.req)

			if tt.req.URL.Path != tt.expectedPath {
				t.Errorf("Expected path %s, got %s", tt.expectedPath, tt.req.URL.Path)
			}

			if tt.req.URL.RawPath != tt.expectedRaw {
				t.Errorf("Expected raw path %s, got %s", tt.expectedRaw, tt.req.URL.RawPath)
			}

			if tt.req.RequestURI != tt.expectedURI {
				t.Errorf("Expected request URI %s, got %s", tt.expectedURI, tt.req.RequestURI)
			}
		})
	}
}
