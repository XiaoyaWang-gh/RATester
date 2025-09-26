package tplimpl

import (
	"fmt"
	"testing"
)

func TestisBaseTemplatePath_1(t *testing.T) {
	tests := []struct {
		name string
		path string
		want bool
	}{
		{
			name: "Test case 1",
			path: "/path/to/base/file.txt",
			want: true,
		},
		{
			name: "Test case 2",
			path: "/path/to/non/base/file.txt",
			want: false,
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

			if got := isBaseTemplatePath(tt.path); got != tt.want {
				t.Errorf("isBaseTemplatePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
