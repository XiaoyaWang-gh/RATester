package rules

import (
	"fmt"
	"testing"
)

func TestNewParser_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	matchers := []string{"matcher1", "matcher2"}
	parser, err := NewParser(matchers)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err = parser.Parse("matcher1(value1, value2) AND matcher2(value3, value4)")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
