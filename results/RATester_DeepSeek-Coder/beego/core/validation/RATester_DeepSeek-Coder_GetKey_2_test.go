package validation

import (
	"fmt"
	"regexp"
	"testing"
)

func TestGetKey_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	email := Email{
		Match: Match{
			Regexp: regexp.MustCompile("^[a-zA-Z0-9.+_-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,15}$"),
		},
		Key: "email",
	}

	expected := "email"
	actual := email.GetKey()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
