package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDenyFragment_1(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	tests := []struct {
		name   string
		url    string
		status int
	}{
		{
			name:   "No fragment",
			url:    "http://example.com/path",
			status: http.StatusOK,
		},
		{
			name:   "With fragment",
			url:    "http://example.com/path#fragment",
			status: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			req, err := http.NewRequest("GET", tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := denyFragment(handler)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.status {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.status)
			}
		})
	}
}
