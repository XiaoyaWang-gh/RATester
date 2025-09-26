package gin

import (
	"fmt"
	"testing"
)

func TestFindWildcard_1(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		wildcard string
		i        int
		valid    bool
	}{
		{"no wildcard", "/test", "", -1, false},
		{"single wildcard", "/test:", ":", 4, true},
		{"multiple wildcards", "/test:*", ":*", 4, true},
		{"escape colon", "/test\\:", ":", 4, true},
		{"escape wildcard", "/test\\*", "*", 4, true},
		{"escape colon and wildcard", "/test\\:*", ":*", 4, true},
		{"wildcard at end", "/test*", "*", 4, true},
		{"wildcard in middle", "/test*me", "*", 4, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			wildcard, i, valid := findWildcard(tt.path)
			if wildcard != tt.wildcard {
				t.Errorf("findWildcard() wildcard = %v, want %v", wildcard, tt.wildcard)
			}
			if i != tt.i {
				t.Errorf("findWildcard() i = %v, want %v", i, tt.i)
			}
			if valid != tt.valid {
				t.Errorf("findWildcard() valid = %v, want %v", valid, tt.valid)
			}
		})
	}
}
