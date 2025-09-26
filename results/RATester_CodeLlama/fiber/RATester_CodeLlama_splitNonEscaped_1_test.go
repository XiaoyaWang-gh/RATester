package fiber

import (
	"fmt"
	"testing"
)

func TestSplitNonEscaped_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := "a,b,c,d,e"
	sep := ","
	result := splitNonEscaped(s, sep)
	if len(result) != 5 {
		t.Errorf("Expected 5 elements, got %d", len(result))
	}
	if result[0] != "a" {
		t.Errorf("Expected 'a', got %s", result[0])
	}
	if result[1] != "b" {
		t.Errorf("Expected 'b', got %s", result[1])
	}
	if result[2] != "c" {
		t.Errorf("Expected 'c', got %s", result[2])
	}
	if result[3] != "d" {
		t.Errorf("Expected 'd', got %s", result[3])
	}
	if result[4] != "e" {
		t.Errorf("Expected 'e', got %s", result[4])
	}
}
