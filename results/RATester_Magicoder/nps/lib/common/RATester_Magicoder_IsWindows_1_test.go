package common

import (
	"fmt"
	"testing"
)

func TestIsWindows_1(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"Test case 1", true},
		{"Test case 2", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := IsWindows(); got != tt.want {
				t.Errorf("IsWindows() = %v, want %v", got, tt.want)
			}
		})
	}
}
