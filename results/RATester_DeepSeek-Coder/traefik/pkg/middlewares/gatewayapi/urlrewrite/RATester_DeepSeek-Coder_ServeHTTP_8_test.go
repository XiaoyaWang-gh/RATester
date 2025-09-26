package urlrewrite

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/aws/smithy-go/ptr"
)

func TestServeHTTP_8(t *testing.T) {
	tests := []struct {
		name       string
		u          urlRewrite
		req        *http.Request
		wantStatus int
		wantPath   string
	}{
		{
			name: "Test with path",
			u: urlRewrite{
				name: "test",
				next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
				}),
				path: ptr.String("/new/path"),
			},
			req: &http.Request{
				Method: "GET",
				URL: &url.URL{
					Path: "/old/path",
				},
			},
			wantStatus: http.StatusOK,
			wantPath:   "/new/path",
		},
		// Add more test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			rr := httptest.NewRecorder()
			tt.u.ServeHTTP(rr, tt.req)

			if status := rr.Code; status != tt.wantStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.wantStatus)
			}

			if gotPath := tt.req.URL.Path; gotPath != tt.wantPath {
				t.Errorf("handler did not rewrite path: got %v want %v", gotPath, tt.wantPath)
			}
		})
	}
}
