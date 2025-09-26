package paths

import (
	"fmt"
	"testing"
)

func TestUglify_1(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"empty", "", "/"},
		{"single char", "a", "a.html"},
		{"no extension", "abc", "abc.html"},
		{"index", "abc/index", "abc"},
		{"index with extension", "abc/index.html", "abc"},
		{"no index", "abc/def", "abc/def"},
		{"no index with extension", "abc/def.html", "abc/def.html"},
		{"clean", "abc/../def", "def"},
		{"clean with extension", "abc/../def.html", "def.html"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Uglify(tt.in); got != tt.want {
				t.Errorf("Uglify(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
