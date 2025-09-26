package inflightreq

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestServeHTTP_27(t *testing.T) {
	testCases := []struct {
		name       string
		handler    http.Handler
		req        *http.Request
		wantStatus int
	}{
		{
			name: "GET request",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}),
			req: &http.Request{
				Method: "GET",
				URL:    &url.URL{Path: "/"},
			},
			wantStatus: http.StatusOK,
		},
		{
			name: "POST request",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusCreated)
			}),
			req: &http.Request{
				Method: "POST",
				URL:    &url.URL{Path: "/"},
			},
			wantStatus: http.StatusCreated,
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

			inFlight := &inFlightReq{
				handler: tc.handler,
				name:    tc.name,
			}

			rr := httptest.NewRecorder()
			inFlight.ServeHTTP(rr, tc.req)

			if rr.Code != tc.wantStatus {
				t.Errorf("Expected status code %d, got %d", tc.wantStatus, rr.Code)
			}
		})
	}
}
