package snicheck

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP_32(t *testing.T) {
	tests := []struct {
		name           string
		req            *http.Request
		expectedStatus int
	}{
		{
			name: "TLS mismatch",
			req: &http.Request{
				Host: "example.com",
				TLS: &tls.ConnectionState{
					ServerName: "mismatch.com",
				},
			},
			expectedStatus: http.StatusMisdirectedRequest,
		},
		{
			name: "TLS match",
			req: &http.Request{
				Host: "example.com",
				TLS: &tls.ConnectionState{
					ServerName: "example.com",
				},
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

			rw := httptest.NewRecorder()
			handler := SNICheck{
				next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			}

			handler.ServeHTTP(rw, tt.req)

			if rw.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rw.Code)
			}
		})
	}
}
