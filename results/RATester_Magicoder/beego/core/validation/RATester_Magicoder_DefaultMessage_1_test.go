package validation

import (
	"fmt"
	"regexp"
	"testing"
)

func TestDefaultMessage_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	regex, _ := regexp.Compile("^[a-zA-Z0-9]+$")
	m := Match{
		Regexp: regex,
		Key:    "testKey",
	}

	expected := fmt.Sprintf(MessageTmpls["Match"], m.Regexp.String())
	actual := m.DefaultMessage()

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
