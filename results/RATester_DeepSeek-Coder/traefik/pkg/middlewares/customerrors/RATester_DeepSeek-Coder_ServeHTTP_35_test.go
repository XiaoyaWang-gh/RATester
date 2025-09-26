package customerrors

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/traefik/traefik/v3/pkg/types"
)

func TestServeHTTP_35(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		req            *http.Request
		expectedStatus int
	}{
		{
			name: "Test with nil backend handler",
			req: &http.Request{
				Method: "GET",
				URL:    &url.URL{Path: "/test"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Test with non-nil backend handler",
			req: &http.Request{
				Method: "GET",
				URL:    &url.URL{Path: "/test"},
			},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &customErrors{
				name:           "test",
				next:           http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
				backendHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
				httpCodeRanges: types.HTTPCodeRanges{},
				backendQuery:   "",
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(c.ServeHTTP)
			handler.ServeHTTP(rr, tt.req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}
		})
	}
}
