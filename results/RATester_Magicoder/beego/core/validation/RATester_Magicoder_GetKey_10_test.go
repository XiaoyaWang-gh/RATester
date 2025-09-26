package validation

import (
	"fmt"
	"regexp"
	"testing"
)

func TestGetKey_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ip := IP{
		Match: Match{
			Regexp: regexp.MustCompile("^[0-9]+\\.[0-9]+\\.[0-9]+\\.[0-9]+$"),
		},
		Key: "ip",
	}

	expected := "ip"
	actual := ip.GetKey()

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
