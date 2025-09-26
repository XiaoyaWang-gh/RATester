package common

import (
	"fmt"
	"testing"
)

func TestGetRunPath_1(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Test case 1",
			want: "path1",
		},
		{
			name: "Test case 2",
			want: "path2",
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

			if got := GetRunPath(); got != tt.want {
				t.Errorf("GetRunPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
