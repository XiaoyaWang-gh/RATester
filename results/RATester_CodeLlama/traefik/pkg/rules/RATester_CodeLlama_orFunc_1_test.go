package rules

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestOrFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	left := func() *Tree {
		return &Tree{
			Matcher: "left",
		}
	}
	right := func() *Tree {
		return &Tree{
			Matcher: "right",
		}
	}
	tree := orFunc(left, right)()
	assert.Equal(t, "or", tree.Matcher)
	assert.Equal(t, "left", tree.RuleLeft.Matcher)
	assert.Equal(t, "right", tree.RuleRight.Matcher)
}
