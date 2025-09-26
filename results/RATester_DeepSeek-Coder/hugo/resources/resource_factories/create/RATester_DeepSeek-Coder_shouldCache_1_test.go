package create

import (
	"fmt"
	"net/http"
	"testing"
)

func TestShouldCache_1(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		want       bool
	}{
		{
			name:       "Test 1: Status OK",
			statusCode: http.StatusOK,
			want:       true,
		},
		{
			name:       "Test 2: Status Moved Permanently",
			statusCode: http.StatusMovedPermanently,
			want:       false,
		},
		{
			name:       "Test 3: Status Found",
			statusCode: http.StatusFound,
			want:       false,
		},
		{
			name:       "Test 4: Status See Other",
			statusCode: http.StatusSeeOther,
			want:       false,
		},
		{
			name:       "Test 5: Status Temporary Redirect",
			statusCode: http.StatusTemporaryRedirect,
			want:       false,
		},
		{
			name:       "Test 6: Status Permanent Redirect",
			statusCode: http.StatusPermanentRedirect,
			want:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := shouldCache(tt.statusCode); got != tt.want {
				t.Errorf("shouldCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
