package validation

import (
	"fmt"
	"regexp"
	"testing"
)

func TestGetKey_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mobile := Mobile{
		Match: Match{
			Regexp: regexp.MustCompile("^[0-9]+$"),
		},
		Key: "1234567890",
	}

	expected := "1234567890"
	actual := mobile.GetKey()

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
