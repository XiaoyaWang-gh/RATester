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
			Regexp: regexp.MustCompile("^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\." +
				"(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\." +
				"(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\." +
				"(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$"),
		},
		Key: "ip",
	}

	expected := "ip"
	actual := ip.GetKey()

	if actual != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}
