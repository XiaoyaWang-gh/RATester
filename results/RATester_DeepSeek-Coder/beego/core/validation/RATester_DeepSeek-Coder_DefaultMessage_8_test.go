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

	m := Mobile{
		Match: Match{
			Regexp: regexp.MustCompile("^1[3-9]{1}\\d{8}$"),
		},
		Key: "Mobile",
	}

	expected := MessageTmpls["Mobile"]
	actual := m.DefaultMessage()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
