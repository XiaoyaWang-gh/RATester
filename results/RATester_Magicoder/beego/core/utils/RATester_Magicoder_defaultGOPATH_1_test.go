package utils

import (
	"fmt"
	"testing"
)

func TestdefaultGOPATH_1(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Test case 1",
			want: "/path/to/go",
		},
		{
			name: "Test case 2",
			want: "/path/to/go",
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

			if got := defaultGOPATH(); got != tt.want {
				t.Errorf("defaultGOPATH() = %v, want %v", got, tt.want)
			}
		})
	}
}
