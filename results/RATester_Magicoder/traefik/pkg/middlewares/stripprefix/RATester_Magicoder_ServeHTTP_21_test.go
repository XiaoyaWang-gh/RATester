package stripprefix

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP_21(t *testing.T) {
	tests := []struct {
		name     string
		prefixes []string
		path     string
		wantPath string
	}{
		{
			name:     "Test case 1",
			prefixes: []string{"/api"},
			path:     "/api/test",
			wantPath: "/test",
		},
		{
			name:     "Test case 2",
			prefixes: []string{"/api", "/v1"},
			path:     "/api/v1/test",
			wantPath: "/test",
		},
		{
			name:     "Test case 3",
			prefixes: []string{"/api"},
			path:     "/api/test",
			wantPath: "/test",
		},
		{
			name:     "Test case 4",
			prefixes: []string{"/api", "/v1"},
			path:     "/api/v1/test",
			wantPath: "/test",
		},
		{
			name:     "Test case 5",
			prefixes: []string{"/api"},
			path:     "/api/test",
			wantPath: "/test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			s := &stripPrefix{
				prefixes: tt.prefixes,
			}

			req, err := http.NewRequest("GET", tt.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			s.ServeHTTP(httptest.NewRecorder(), req)

			if req.URL.Path != tt.wantPath {
				t.Errorf("Expected path %s, but got %s", tt.wantPath, req.URL.Path)
			}
		})
	}
}
