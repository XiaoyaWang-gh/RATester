package glob

import (
	"fmt"
	"testing"

	"github.com/gobwas/glob"
)

func TestMatch_4(t *testing.T) {
	tests := []struct {
		name      string
		isWindows bool
		pattern   string
		input     string
		want      bool
	}{
		{
			name:      "Match on Windows",
			isWindows: true,
			pattern:   "*.txt",
			input:     "test.txt",
			want:      true,
		},
		{
			name:      "No Match on Windows",
			isWindows: true,
			pattern:   "*.txt",
			input:     "test.doc",
			want:      false,
		},
		{
			name:      "Match on Unix",
			isWindows: false,
			pattern:   "*.txt",
			input:     "test.txt",
			want:      true,
		},
		{
			name:      "No Match on Unix",
			isWindows: false,
			pattern:   "*.txt",
			input:     "test.doc",
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			g := globDecorator{
				isWindows: tt.isWindows,
				g:         glob.MustCompile(tt.pattern),
			}

			if got := g.Match(tt.input); got != tt.want {
				t.Errorf("Match() = %v, want %v", got, tt.want)
			}
		})
	}
}
