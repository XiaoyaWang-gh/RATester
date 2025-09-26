package langs

import (
	"fmt"
	"testing"

	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

func TestCompareStrings_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Collator{
		c: collate.New(language.Und, collate.Loose),
	}

	a := "a"
	b := "b"

	if c.CompareStrings(a, b) != -1 {
		t.Errorf("Expected %q to be less than %q", a, b)
	}

	a = "b"
	b = "a"

	if c.CompareStrings(a, b) != 1 {
		t.Errorf("Expected %q to be greater than %q", a, b)
	}

	a = "a"
	b = "a"

	if c.CompareStrings(a, b) != 0 {
		t.Errorf("Expected %q to be equal to %q", a, b)
	}
}
