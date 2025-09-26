package service

import (
	"fmt"
	"testing"
)

func TeststatusText_1(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		want       string
	}{
		{
			name:       "Test Case 1",
			statusCode: 200,
			want:       "OK",
		},
		{
			name:       "Test Case 2",
			statusCode: 404,
			want:       "Not Found",
		},
		{
			name:       "Test Case 3",
			statusCode: 500,
			want:       "Internal Server Error",
		},
		{
			name:       "Test Case 4",
			statusCode: 499,
			want:       "Client Closed Request",
		},
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
