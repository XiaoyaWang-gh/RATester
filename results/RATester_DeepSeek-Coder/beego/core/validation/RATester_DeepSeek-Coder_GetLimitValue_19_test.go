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

	zip := ZipCode{
		Match: Match{
			Regexp: regexp.MustCompile("^[0-9]{5}(?:-[0-9]{4})?$"),
		},
		Key: "zip",
	}

	result := zip.GetLimitValue()

	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}
}
