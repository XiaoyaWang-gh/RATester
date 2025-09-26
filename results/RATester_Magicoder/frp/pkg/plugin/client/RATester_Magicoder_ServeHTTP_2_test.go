package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP_2(t *testing.T) {
	tests := []struct {
		name   string
		method string
		want   int
	}{
		{
			name:   "Test GET request",
			method: http.MethodGet,
			want:   http.StatusOK,
		},
		{
			name:   "Test POST request",
			method: http.MethodPost,
			want:   http.StatusOK,
		},
		{
			name:   "Test CONNECT request",
			method: http.MethodConnect,
			want:   http.StatusOK,
		},
		{
			name:   "Test UNKNOWN request",
			method: "UNKNOWN",
			want:   http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			req, err := http.NewRequest(tt.method, "/", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := &HTTPProxy{}

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.want {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.want)
			}
		})
	}
}
