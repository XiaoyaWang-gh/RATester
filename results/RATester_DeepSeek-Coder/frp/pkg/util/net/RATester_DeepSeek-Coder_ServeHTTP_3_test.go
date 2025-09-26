package net

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP_3(t *testing.T) {
	testCases := []struct {
		name           string
		request        *http.Request
		expectedHeader string
	}{
		{
			name: "Test with Accept-Encoding header containing gzip",
			request: &http.Request{
				Header: http.Header{
					"Accept-Encoding": []string{"gzip"},
				},
			},
			expectedHeader: "gzip",
		},
		{
			name: "Test with Accept-Encoding header not containing gzip",
			request: &http.Request{
				Header: http.Header{
					"Accept-Encoding": []string{"deflate"},
				},
			},
			expectedHeader: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gw := &HTTPGzipWrapper{
				h: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
				}),
			}

			rr := httptest.NewRecorder()
			gw.ServeHTTP(rr, tc.request)

			result := rr.Header().Get("Content-Encoding")
			if result != tc.expectedHeader {
				t.Errorf("Expected header 'Content-Encoding' to be '%s', but got '%s'", tc.expectedHeader, result)
			}
		})
	}
}
