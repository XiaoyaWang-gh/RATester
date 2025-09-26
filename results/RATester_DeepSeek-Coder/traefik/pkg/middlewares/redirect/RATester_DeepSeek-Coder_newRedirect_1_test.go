package redirect

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNewRedirect_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		next        http.Handler
		regex       string
		replacement string
		permanent   bool
		rawURL      func(*http.Request) string
		wantErr     bool
	}{
		{
			name:        "valid",
			next:        http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			regex:       "^/old/(.*)$",
			replacement: "/new/$1",
			permanent:   true,
			rawURL:      func(r *http.Request) string { return r.URL.String() },
			wantErr:     false,
		},
		{
			name:        "invalid regex",
			next:        http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			regex:       "^/old/(.*",
			replacement: "/new/$1",
			permanent:   true,
			rawURL:      func(r *http.Request) string { return r.URL.String() },
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, err := newRedirect(tt.next, tt.regex, tt.replacement, tt.permanent, tt.rawURL, "")
			if (err != nil) != tt.wantErr {
				t.Errorf("newRedirect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
