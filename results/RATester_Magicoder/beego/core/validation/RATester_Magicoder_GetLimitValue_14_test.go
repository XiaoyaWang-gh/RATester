package validation

import (
	"fmt"
	"regexp"
	"testing"
)

func TestGetLimitValue_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	regex, _ := regexp.Compile("^[a-zA-Z0-9]+$")
	match := Match{
		Regexp: regex,
		Key:    "testKey",
	}

	expected := regex.String()
	actual := match.GetLimitValue()

	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
