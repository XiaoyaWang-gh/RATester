package security

import (
	"fmt"
	"regexp"
	"testing"
)

func TestString_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := Whitelist{
		patterns: []*regexp.Regexp{
			regexp.MustCompile("^a"),
			regexp.MustCompile("^b"),
		},
		patternsStrings: []string{"^a", "^b"},
	}

	if got := w.String(); got != "[\n^a\n^b\n]" {
		t.Errorf("String() = %v, want %v", got, "[\n^a\n^b\n]")
	}
}
