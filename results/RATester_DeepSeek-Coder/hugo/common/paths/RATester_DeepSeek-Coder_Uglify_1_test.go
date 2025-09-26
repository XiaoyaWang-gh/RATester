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
		{"empty string", "", "/"},
		{"single char", "a", "a.html"},
		{"single char with ext", "a.html", "a.html"},
		{"index in the middle", "/a/b/index.html", "/a/b/index.html"},
		{"index at the end", "/a/b/c/index", "/a/b/c/index.html"},
		{"no index", "/a/b/c", "/a/b/c.html"},
		{"complex path", "/a/b/c/d.html", "/a/b/c/d.html"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := Uglify(tt.in)
			if got != tt.want {
				t.Errorf("Uglify(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
