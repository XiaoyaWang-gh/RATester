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
			Regexp: regexp.MustCompile("^[a-z]+$"),
		},
		Key: "key",
	}
	if n.GetLimitValue() != n.Regexp.String() {
		t.Errorf("GetLimitValue() = %v, want %v", n.GetLimitValue(), n.Regexp.String())
	}
}
