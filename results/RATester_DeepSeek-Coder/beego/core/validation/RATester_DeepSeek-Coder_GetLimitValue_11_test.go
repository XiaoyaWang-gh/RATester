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
			Regexp: regexp.MustCompile("^[0-9]+$"),
		},
		Key: "tel",
	}

	limitValue := testTel.GetLimitValue()
	if limitValue != nil {
		t.Errorf("Expected nil, got %v", limitValue)
	}
}
