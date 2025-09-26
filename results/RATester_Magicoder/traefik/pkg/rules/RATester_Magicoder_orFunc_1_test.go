package rules

import (
	"fmt"
	"testing"
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
	result := orFunc(left, right)()
	if result.Matcher != "or" {
		t.Errorf("Expected matcher to be 'or', but got %s", result.Matcher)
	}
	if result.RuleLeft.Matcher != "left" {
		t.Errorf("Expected left rule matcher to be 'left', but got %s", result.RuleLeft.Matcher)
	}
	if result.RuleRight.Matcher != "right" {
		t.Errorf("Expected right rule matcher to be 'right', but got %s", result.RuleRight.Matcher)
	}
}
