package validation

import (
	"fmt"
	"regexp"
	"testing"
)

func TestGetLimitValue_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testTel := Tel{
		Match: Match{
			Regexp: regexp.MustCompile("^[0-9]{11}$"),
		},
		Key: "phone",
	}

	result := testTel.GetLimitValue()

	if result != nil {
		t.Errorf("Expected nil, but got %v", result)
	}
}
