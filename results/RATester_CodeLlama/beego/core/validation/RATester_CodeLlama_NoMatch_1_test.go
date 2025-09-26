package validation

import (
	"fmt"
	"regexp"
	"testing"
)

func TestNoMatch_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := &Validation{}
	obj := "1234567890"
	regex := regexp.MustCompile("^[0-9]{10}$")
	key := "phone"
	result := v.NoMatch(obj, regex, key)
	if result.Error == nil {
		t.Errorf("NoMatch failed for %s", key)
	}
}
