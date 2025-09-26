package rules

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestNotFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	elem := func() *Tree {
		return &Tree{
			Matcher: "foo",
		}
	}
	expected := &Tree{
		Matcher: "foo",
		Not:     true,
	}

	// when
	actual := notFunc(elem)

	// then
	assert.Equal(t, expected, actual())
}
