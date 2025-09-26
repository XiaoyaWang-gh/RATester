package validation

import (
	"fmt"
	"regexp"
	"testing"
)

func TestDefaultMessage_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := Email{
		Match: Match{
			Regexp: regexp.MustCompile("^[a-zA-Z0-9.%+\\-]+@[a-zA-Z0-9.\\-]+\\.[a-zA-Z]{2,}$"),
		},
		Key: "email",
	}

	expected := "Email"
	actual := e.DefaultMessage()

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
