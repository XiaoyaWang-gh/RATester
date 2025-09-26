package validation

import (
	"fmt"
	"regexp"
	"testing"
)

func TestGetLimitValue_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ip := IP{
		Match: Match{
			Regexp: regexp.MustCompile("^[0-9]+$"),
		},
		Key: "testKey",
	}

	result := ip.GetLimitValue()

	if result != nil {
		t.Errorf("Expected nil, but got %v", result)
	}
}
