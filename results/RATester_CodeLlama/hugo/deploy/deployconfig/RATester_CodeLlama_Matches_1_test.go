package deployconfig

import (
	"fmt"
	"regexp"
	"testing"
)

func TestMatches_1(t *testing.T) {
	tests := []struct {
		name string
		m    *Matcher
		path string
		want bool
	}{
		{
			name: "Matches",
			m: &Matcher{
				Pattern: "*.html",
				Re:      regexp.MustCompile("^.*\\.html$"),
			},
			path: "index.html",
			want: true,
		},
		{
			name: "DoesNotMatch",
			m: &Matcher{
				Pattern: "*.html",
				Re:      regexp.MustCompile("^.*\\.html$"),
			},
			path: "index.htm",
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

			if got := tt.m.Matches(tt.path); got != tt.want {
				t.Errorf("Matches() = %v, want %v", got, tt.want)
			}
		})
	}
}
