package validation

import (
	"fmt"
	"regexp"
	"testing"
)

func TestDefaultMessage_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := NoMatch{
		Match: Match{
			Regexp: regexp.MustCompile("^[a-zA-Z0-9]+$"),
		},
		Key: "testKey",
	}

	expected := fmt.Sprintf(MessageTmpls["NoMatch"], n.Regexp.String())
	actual := n.DefaultMessage()

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
