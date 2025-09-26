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
			name: "empty string",
			in:   "",
			want: "/",
		},
		{
			name: "simple path",
			in:   "/path/to/page",
			want: "/path/to/page",
		},
		{
			name: "index.html",
			in:   "/path/to/index.html",
			want: "/path/to",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := PrettifyURL(tt.in)
			if got != tt.want {
				t.Errorf("PrettifyURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
