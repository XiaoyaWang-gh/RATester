package validation

import (
	"fmt"
	"regexp"
	"testing"
)

func TestDefaultMessage_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := Base64{
		Match: Match{
			Regexp: regexp.MustCompile("^[A-Za-z0-9+/]*={0,2}$"),
		},
		Key: "Base64",
	}

	expected := "Base64"
	actual := b.DefaultMessage()

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
