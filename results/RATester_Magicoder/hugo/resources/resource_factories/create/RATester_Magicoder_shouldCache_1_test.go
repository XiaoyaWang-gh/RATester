package create

import (
	"fmt"
	"net/http"
	"testing"
)

func TestshouldCache_1(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		want       bool
	}{
		{
			name:       "Test case 1: StatusMovedPermanently",
			statusCode: http.StatusMovedPermanently,
			want:       false,
		},
		{
			name:       "Test case 2: StatusFound",
			statusCode: http.StatusFound,
			want:       false,
		},
		{
			name:       "Test case 3: StatusSeeOther",
			statusCode: http.StatusSeeOther,
			want:       false,
		},
		{
			name:       "Test case 4: StatusTemporaryRedirect",
			statusCode: http.StatusTemporaryRedirect,
			want:       false,
		},
		{
			name:       "Test case 5: StatusPermanentRedirect",
			statusCode: http.StatusPermanentRedirect,
			want:       false,
		},
		{
			name:       "Test case 6: StatusOK",
			statusCode: http.StatusOK,
			want:       true,
		},
		{
			name:       "Test case 7: StatusBadRequest",
			statusCode: http.StatusBadRequest,
			want:       true,
		},
		{
			name:       "Test case 8: StatusNotFound",
			statusCode: http.StatusNotFound,
			want:       true,
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
