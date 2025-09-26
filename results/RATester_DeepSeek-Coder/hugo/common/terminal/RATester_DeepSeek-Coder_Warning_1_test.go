package terminal

import (
	"fmt"
	"testing"
)

func TestWarning_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	expected := "\033[33mWarning: Hello\033[0m"
	result := Warning("Hello")
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
