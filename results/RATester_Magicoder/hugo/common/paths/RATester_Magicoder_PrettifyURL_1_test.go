package paths

import (
	"fmt"
	"testing"
)

func TestPrettifyURL_1(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "Test case 1",
			in:   "/path/to/file.html",
			want: "/path/to",
		},
		{
			name: "Test case 2",
			in:   "/path/to/index.html",
			want: "/path/to",
		},
		{
			name: "Test case 3",
			in:   "",
			want: "/",
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

			if got := PrettifyURL(tt.in); got != tt.want {
				t.Errorf("PrettifyURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
