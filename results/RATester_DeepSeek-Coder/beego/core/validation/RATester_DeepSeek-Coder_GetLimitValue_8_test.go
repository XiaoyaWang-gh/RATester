package validation

import (
	"fmt"
	"regexp"
	"testing"
)

func TestGetLimitValue_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	email := Email{
		Match: Match{
			Regexp: regexp.MustCompile("^[a-zA-Z0-9.+_-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,15}$"),
		},
		Key: "email",
	}

	result := email.GetLimitValue()

	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}
}
