package cors

import (
	"fmt"
	"testing"
)

func TestMatchScheme_1(t *testing.T) {
	type test struct {
		domain  string
		pattern string
		want    bool
	}

	tests := []test{
		{"http:", "http:", true},
		{"https:", "https:", true},
		{"http:", "https:", false},
		{"https:", "http:", false},
		{"http:", ":", false},
		{":", "http:", false},
		{"", "", true},
		{"http:", "", false},
		{"", "http:", false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("domain: %s, pattern: %s", tt.domain, tt.pattern), func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := matchScheme(tt.domain, tt.pattern)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
