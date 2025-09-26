package validation

import (
	"fmt"
	"regexp"
	"testing"

	"gotest.tools/assert"
)

func TestDefaultMessage_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := Match{
		Regexp: regexp.MustCompile("^[a-z]+$"),
		Key:    "name",
	}
	assert.Equal(t, "name must match ^[a-z]+$", m.DefaultMessage())
}
