package validation

import (
	"fmt"
	"regexp"
	"testing"
)

func TestDefaultMessage_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mobile := Mobile{
		Match: Match{
			Regexp: regexp.MustCompile("^1[34578][0-9]{9}$"),
		},
		Key: "Mobile",
	}

	expected := "Mobile"
	actual := mobile.DefaultMessage()

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
