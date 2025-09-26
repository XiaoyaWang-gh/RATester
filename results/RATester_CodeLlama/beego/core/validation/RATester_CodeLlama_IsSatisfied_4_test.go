package validation

import (
	"fmt"
	"regexp"
	"testing"
)

func TestIsSatisfied_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := NoMatch{
		Match: Match{
			Regexp: regexp.MustCompile("^[a-z]+$"),
		},
		Key: "key",
	}
	obj := "123"
	if n.IsSatisfied(obj) {
		t.Errorf("Expected %v to not match %v", obj, n.Match.Regexp)
	}
}
