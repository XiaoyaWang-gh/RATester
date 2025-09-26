package logs

import (
	"fmt"
	"testing"
)

func TestResetColor_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	expected := "\033[0m"
	result := ResetColor()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
