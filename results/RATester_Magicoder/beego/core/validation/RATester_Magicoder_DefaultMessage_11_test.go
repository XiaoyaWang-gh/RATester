package validation

import (
	"fmt"
	"regexp"
	"testing"
)

func TestDefaultMessage_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	z := ZipCode{
		Match: Match{
			Regexp: regexp.MustCompile("^[0-9]{5}(?:-[0-9]{4})?$"),
		},
		Key: "ZipCode",
	}

	expected := "ZipCode is not valid"
	actual := z.DefaultMessage()

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
