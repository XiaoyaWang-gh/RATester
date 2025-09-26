package paths

import (
	"fmt"
	"testing"
)

func TestPrettifyURLPath_1(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "Test case 1",
			in:   "/path/to/resource",
			want: "/path/to/resource",
		},
		{
			name: "Test case 2",
			in:   "/path/to/resource/",
			want: "/path/to/resource",
		},
		{
			name: "Test case 3",
			in:   "/path/to/resource//",
			want: "/path/to/resource",
		},
		{
			name: "Test case 4",
			in:   "/path/to/resource///",
			want: "/path/to/resource",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := PrettifyURLPath(tt.in); got != tt.want {
				t.Errorf("PrettifyURLPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
