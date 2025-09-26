package validation

import (
	"fmt"
	"regexp"
	"testing"
)

func TestGetLimitValue_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := Base64{
		Match: Match{
			Regexp: regexp.MustCompile("^[A-Za-z0-9+/]*={0,2}$"),
		},
		Key: "testKey",
	}

	result := b.GetLimitValue()

	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}
}
