package helpers

import (
	"fmt"
	"testing"
)

func TestcacheDirDefault_1(t *testing.T) {

	tests := []struct {
		name     string
		cacheDir string
		want     string
	}{
		{
			name:     "Test case 1",
			cacheDir: "/path/to/cache",
			want:     "/path/to/cache/",
		},
		{
			name:     "Test case 2",
			cacheDir: "",
			want:     "",
		},
		{
			name:     "Test case 3",
			cacheDir: "/path/to/cache/",
			want:     "/path/to/cache/",
		},
		{
			name:     "Test case 4",
			cacheDir: "/path/to/cache/subdir",
			want:     "/path/to/cache/subdir/",
		},
		{
			name:     "Test case 5",
			cacheDir: "/path/to/cache/subdir/",
			want:     "/path/to/cache/subdir/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := cacheDirDefault(tt.cacheDir); got != tt.want {
				t.Errorf("cacheDirDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
