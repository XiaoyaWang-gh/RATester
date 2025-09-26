package security

import (
	"fmt"
	"testing"
)

func TestString_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := Whitelist{
		patternsStrings: []string{"pattern1", "pattern2", "pattern3"},
	}

	expected := "[pattern1 pattern2 pattern3]"
	result := w.String()

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
