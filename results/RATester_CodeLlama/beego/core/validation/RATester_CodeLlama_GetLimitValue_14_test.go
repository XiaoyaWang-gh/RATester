package validation

import (
	"fmt"
	"regexp"
	"testing"
)

func TestGetLimitValue_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := Match{
		Regexp: regexp.MustCompile("^[a-z]+$"),
		Key:    "name",
	}
	if m.GetLimitValue() != m.Regexp.String() {
		t.Errorf("GetLimitValue() = %v, want %v", m.GetLimitValue(), m.Regexp.String())
	}
}
