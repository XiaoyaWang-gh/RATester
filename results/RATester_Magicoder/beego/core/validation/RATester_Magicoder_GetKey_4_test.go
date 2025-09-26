package validation

import (
	"fmt"
	"regexp"
	"testing"
)

func TestGetKey_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	match := Match{
		Regexp: regexp.MustCompile("^[a-zA-Z0-9]+$"),
		Key:    "testKey",
	}

	key := match.GetKey()

	if key != "testKey" {
		t.Errorf("Expected 'testKey', but got '%s'", key)
	}
}
