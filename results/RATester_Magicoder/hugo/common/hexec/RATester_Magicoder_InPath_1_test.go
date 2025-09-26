package hexec

import (
	"fmt"
	"testing"
)

func TestInPath_1(t *testing.T) {
	tests := []struct {
		name       string
		binaryName string
		want       bool
	}{
		{
			name:       "Test case 1: Binary exists",
			binaryName: "ls",
			want:       true,
		},
		{
			name:       "Test case 2: Binary does not exist",
			binaryName: "nonexistentbinary",
			want:       false,
		},
		{
			name:       "Test case 3: Binary name contains slash",
			binaryName: "ls/",
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

			if got := InPath(tt.binaryName); got != tt.want {
				t.Errorf("InPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
