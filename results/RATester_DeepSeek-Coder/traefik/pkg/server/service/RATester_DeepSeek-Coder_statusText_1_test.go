package service

import (
	"fmt"
	"net/http"
	"testing"
)

func TestStatusText_1(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		want       string
	}{
		{
			name:       "StatusClientClosedRequest",
			statusCode: StatusClientClosedRequest,
			want:       StatusClientClosedRequestText,
		},
		{
			name:       "StatusOK",
			statusCode: http.StatusOK,
			want:       http.StatusText(http.StatusOK),
		},
		{
			name:       "StatusNotFound",
			statusCode: http.StatusNotFound,
			want:       http.StatusText(http.StatusNotFound),
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := statusText(tt.statusCode); got != tt.want {
				t.Errorf("statusText() = %v, want %v", got, tt.want)
			}
		})
	}
}
