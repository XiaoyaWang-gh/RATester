package validation

import (
	"fmt"
	"regexp"
	"testing"
)

func TestMatch_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := &Validation{}
	obj := "1234567890"
	regex := regexp.MustCompile("^[0-9]{10}$")
	key := "phone"
	result := v.Match(obj, regex, key)
	if result.Error != nil {
		t.Errorf("Match failed: %s", result.Error.Message)
	}
}
