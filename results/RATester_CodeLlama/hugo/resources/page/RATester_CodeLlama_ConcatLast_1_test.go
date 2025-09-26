package page

import (
	"fmt"
	"testing"
)

func TestConcatLast_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pagePathBuilder{}
	p.Add("a")
	p.ConcatLast("b")
	if p.Last() != "ab" {
		t.Errorf("Expected 'ab', got '%s'", p.Last())
	}
	p.ConcatLast("c")
	if p.Last() != "abc" {
		t.Errorf("Expected 'abc', got '%s'", p.Last())
	}
	p.ConcatLast("")
	if p.Last() != "abc" {
		t.Errorf("Expected 'abc', got '%s'", p.Last())
	}
	p.ConcatLast("d")
	if p.Last() != "abcd" {
		t.Errorf("Expected 'abcd', got '%s'", p.Last())
	}
}
