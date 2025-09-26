package rules

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseMatchers_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{
		Matcher: "and",
		RuleLeft: &Tree{
			Matcher: "matcher1",
			Value:   []string{"value1"},
		},
		RuleRight: &Tree{
			Matcher: "matcher2",
			Value:   []string{"value2"},
		},
	}

	matchers := []string{"matcher1", "matcher2"}
	expected := []string{"value1", "value2"}

	result := tree.ParseMatchers(matchers)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
