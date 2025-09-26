package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHTTPHandler_1(t *testing.T) {
	testCases := []struct {
		name           string
		request        *http.Request
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Success",
			request: &http.Request{
				Method: "GET",
				URL:    &url.URL{Path: "/success"},
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "Successful response",
		},
		{
			name: "NotFound",
			request: &http.Request{
				Method: "GET",
				URL:    &url.URL{Path: "/not-found"},
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not Found",
		},
		{
			name: "InternalError",
			request: &http.Request{
				Method: "GET",
				URL:    &url.URL{Path: "/internal-error"},
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Internal Server Error",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			rw := httptest.NewRecorder()
			hp := &HTTPProxy{}
			hp.HTTPHandler(rw, tc.request)

			if rw.Code != tc.expectedStatus {
				t.Errorf("Expected status %d, but got %d", tc.expectedStatus, rw.Code)
			}

			if rw.Body.String() != tc.expectedBody {
				t.Errorf("Expected body %q, but got %q", tc.expectedBody, rw.Body.String())
			}
		})
	}
}
