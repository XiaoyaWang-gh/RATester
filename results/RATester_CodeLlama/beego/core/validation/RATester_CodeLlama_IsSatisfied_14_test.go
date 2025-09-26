package validation

import (
	"fmt"
	"regexp"
	"testing"
)

func TestIsSatisfied_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := Match{
		Regexp: regexp.MustCompile("^[a-z]+$"),
		Key:    "name",
	}
	obj := "hello"
	if !m.IsSatisfied(obj) {
		t.Errorf("IsSatisfied() = %v, want %v", false, true)
	}
}
