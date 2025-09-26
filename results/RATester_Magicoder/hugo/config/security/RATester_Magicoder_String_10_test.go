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
		acceptNone:      true,
		patterns:        []*regexp.Regexp{regexp.MustCompile(".*")},
		patternsStrings: []string{"*"},
	}

	expected := "[*]"
	result := w.String()

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
