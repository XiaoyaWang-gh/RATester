package validation

import (
	"fmt"
	"regexp"
	"testing"
)

func TestGetLimitValue_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mobile := Mobile{
		Match: Match{
			Regexp: regexp.MustCompile("^1[34578][0-9]{9}$"),
		},
		Key: "mobile",
	}

	result := mobile.GetLimitValue()

	if result != nil {
		t.Errorf("Expected nil, but got %v", result)
	}
}
