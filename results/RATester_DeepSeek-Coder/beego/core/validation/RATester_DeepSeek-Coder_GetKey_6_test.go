package validation

import (
	"fmt"
	"regexp"
	"testing"
)

func TestGetKey_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := NoMatch{
		Match: Match{
			Regexp: regexp.MustCompile("^[a-zA-Z0-9]*$"),
		},
		Key: "testKey",
	}

	expected := "testKey"
	actual := n.GetKey()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
