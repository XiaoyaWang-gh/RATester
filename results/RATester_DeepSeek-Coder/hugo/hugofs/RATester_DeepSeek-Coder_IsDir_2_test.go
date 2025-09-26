package hugofs

import (
	"fmt"
	"testing"
	"time"
)

func TestIsDir_2(t *testing.T) {
	testCases := []struct {
		name     string
		fileInfo *dirNameOnlyFileInfo
		want     bool
	}{
		{
			name: "TestIsDir_ReturnsTrue",
			fileInfo: &dirNameOnlyFileInfo{
				name:    "test",
				modTime: time.Now(),
			},
			want: true,
		},
		{
			name: "TestIsDir_ReturnsFalse",
			fileInfo: &dirNameOnlyFileInfo{
				name:    "test",
				modTime: time.Now(),
			},
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tc.fileInfo.IsDir()
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
