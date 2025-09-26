package compare

import (
	"fmt"
	"testing"
)

func TestLessStrings_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tests := []struct {
		s, t string
		want bool
	}{
		{"abc", "def", true},
		{"def", "abc", false},
		{"abc", "abc", false},
	}

	for _, test := range tests {
		if got := LessStrings(test.s, test.t); got != test.want {
			t.Errorf("LessStrings(%q, %q) = %v", test.s, test.t, got)
		}
	}
}
