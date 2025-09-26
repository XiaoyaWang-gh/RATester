package testenv

import (
	"fmt"
	"testing"
)

func TestHasGoBuild_1(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "Test case 1",
			want: false,
		},
		{
			name: "Test case 2",
			want: false,
		},
		{
			name: "Test case 3",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := HasGoBuild(); got != tt.want {
				t.Errorf("HasGoBuild() = %v, want %v", got, tt.want)
			}
		})
	}
}
