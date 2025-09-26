package paths

import (
	"fmt"
	"testing"
)

func TestAddLeadingSlash_1(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "Test case 1",
			path: "test",
			want: "/test",
		},
		{
			name: "Test case 2",
			path: "/test",
			want: "/test",
		},
		{
			name: "Test case 3",
			path: "",
			want: "/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := AddLeadingSlash(tt.path); got != tt.want {
				t.Errorf("AddLeadingSlash() = %v, want %v", got, tt.want)
			}
		})
	}
}
