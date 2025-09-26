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
			name: "GET request",
			request: &http.Request{
				Method: "GET",
				URL:    &url.URL{Path: "/test"},
				Header: make(http.Header),
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "expected body",
		},
		{
			name: "POST request",
			request: &http.Request{
				Method: "POST",
				URL:    &url.URL{Path: "/test"},
				Header: make(http.Header),
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "expected body",
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			recorder := httptest.NewRecorder()
			proxy := &HTTPProxy{}

			proxy.HTTPHandler(recorder, tc.request)

			if recorder.Code != tc.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tc.expectedStatus, recorder.Code)
			}

			if recorder.Body.String() != tc.expectedBody {
				t.Errorf("Expected body '%s', got '%s'", tc.expectedBody, recorder.Body.String())
			}
		})
	}
}
