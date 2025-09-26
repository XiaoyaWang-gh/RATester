package validation

import (
	"fmt"
	"regexp"
	"testing"
)

func TestDefaultMessage_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testTel := Tel{
		Match: Match{
			Regexp: regexp.MustCompile(`^1[3456789]\d{9}$`),
		},
		Key: "Tel",
	}

	expected := "Tel"
	actual := testTel.DefaultMessage()

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
