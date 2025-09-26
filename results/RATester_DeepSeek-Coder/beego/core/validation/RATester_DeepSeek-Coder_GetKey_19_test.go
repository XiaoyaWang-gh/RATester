package validation

import (
	"fmt"
	"regexp"
	"testing"
)

func TestGetKey_19(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	zip := ZipCode{
		Match: Match{
			Regexp: regexp.MustCompile("^[0-9]{5}(?:-[0-9]{4})?$"),
		},
		Key: "12345",
	}

	expected := "12345"
	actual := zip.GetKey()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
