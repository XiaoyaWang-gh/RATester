package testenv

import (
	"fmt"
	"testing"
)

func TestHasGoRun_1(t *testing.T) {

	tests := []struct {
		name string
		want bool
	}{
		{"Test case 1", true},
		{"Test case 2", false},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := HasGoRun(); got != tt.want {
				t.Errorf("HasGoRun() = %v, want %v", got, tt.want)
			}
		})
	}
}
