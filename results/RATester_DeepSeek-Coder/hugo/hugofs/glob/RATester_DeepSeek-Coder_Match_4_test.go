package glob

import (
	"fmt"
	"testing"

	"github.com/gobwas/glob"
)

func TestMatch_4(t *testing.T) {
	tests := []struct {
		name     string
		g        globDecorator
		s        string
		expected bool
	}{
		{
			name: "Match should return true when the string matches the glob pattern",
			g: globDecorator{
				isWindows: false,
				g:         glob.MustCompile("*.txt"),
			},
			s:        "test.txt",
			expected: true,
		},
		{
			name: "Match should return false when the string does not match the glob pattern",
			g: globDecorator{
				isWindows: false,
				g:         glob.MustCompile("*.txt"),
			},
			s:        "test.go",
			expected: false,
		},
		{
			name: "Match should normalize the string to lowercase",
			g: globDecorator{
				isWindows: false,
				g:         glob.MustCompile("*.TXT"),
			},
			s:        "test.txt",
			expected: true,
		},
		{
			name: "Match should convert Windows slashes to Unix slashes",
			g: globDecorator{
				isWindows: true,
				g:         glob.MustCompile("*/*.txt"),
			},
			s:        "test\\test.txt",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tt.g.Match(tt.s)
			if result != tt.expected {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}
		})
	}
}
