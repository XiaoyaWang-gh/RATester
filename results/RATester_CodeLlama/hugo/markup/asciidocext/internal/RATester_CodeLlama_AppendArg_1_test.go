package internal

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestAppendArg_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	a := &AsciidocConverter{}
	args := []string{}
	option := "option"
	value := "value"
	defaultValue := "defaultValue"
	allowedValues := map[string]bool{"value": true}

	// when
	args = a.AppendArg(args, option, value, defaultValue, allowedValues)

	// then
	assert.Equal(t, []string{"option", "value"}, args)
}
