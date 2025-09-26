package rules

import (
	"fmt"
	"testing"
)

func TestAndFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	left := func() *Tree {
		return &Tree{
			Matcher: "matcher1",
			Value:   []string{"value1", "value2"},
		}
	}

	right := func() *Tree {
		return &Tree{
			Matcher: "matcher2",
			Value:   []string{"value3", "value4"},
		}
	}

	result := andFunc(left, right)()

	if result.Matcher != "and" {
		t.Errorf("Expected Matcher to be 'and', got %s", result.Matcher)
	}

	if result.RuleLeft.Matcher != "matcher1" {
		t.Errorf("Expected RuleLeft.Matcher to be 'matcher1', got %s", result.RuleLeft.Matcher)
	}

	if len(result.RuleLeft.Value) != 2 || result.RuleLeft.Value[0] != "value1" || result.RuleLeft.Value[1] != "value2" {
		t.Errorf("Expected RuleLeft.Value to be ['value1', 'value2'], got %v", result.RuleLeft.Value)
	}

	if result.RuleRight.Matcher != "matcher2" {
		t.Errorf("Expected RuleRight.Matcher to be 'matcher2', got %s", result.RuleRight.Matcher)
	}

	if len(result.RuleRight.Value) != 2 || result.RuleRight.Value[0] != "value3" || result.RuleRight.Value[1] != "value4" {
		t.Errorf("Expected RuleRight.Value to be ['value3', 'value4'], got %v", result.RuleRight.Value)
	}
}
