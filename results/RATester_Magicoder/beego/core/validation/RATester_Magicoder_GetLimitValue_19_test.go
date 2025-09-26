package validation

import (
	"fmt"
	"regexp"
	"testing"
)

func TestGetLimitValue_19(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	z := ZipCode{
		Match: Match{
			Regexp: regexp.MustCompile(`^[0-9]{5}(?:-[0-9]{4})?$`),
		},
		Key: "zip",
	}

	result := z.GetLimitValue()

	if result != nil {
		t.Errorf("Expected nil, but got %v", result)
	}
}
