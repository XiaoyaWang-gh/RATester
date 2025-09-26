package validation

import (
	"fmt"
	"regexp"
	"testing"
)

func TestGetLimitValue_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := NoMatch{
		Match: Match{
			Regexp: regexp.MustCompile("^[a-zA-Z0-9]+$"),
		},
		Key: "testKey",
	}

	expected := n.Regexp.String()
	actual := n.GetLimitValue()

	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
