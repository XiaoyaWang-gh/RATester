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
			Value:   []string{"value1", "value2"},
		},
		RuleRight: &Tree{
			Matcher: "matcher2",
			Value:   []string{"value3", "value4"},
		},
	}

	matchers := []string{"matcher1", "matcher2"}
	expected := []string{"value1", "value2", "value3", "value4"}
	result := tree.ParseMatchers(matchers)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
