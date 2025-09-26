package validation

import (
	"fmt"
	"regexp"
	"testing"
)

func TestDefaultMessage_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := NoMatch{
		Match: Match{
			Regexp: regexp.MustCompile("^[a-z]+$"),
		},
		Key: "key",
	}
	if n.DefaultMessage() != "key must not match ^[a-z]+$" {
		t.Errorf("DefaultMessage() = %v, want %v", n.DefaultMessage(), "key must not match ^[a-z]+$")
	}
}
