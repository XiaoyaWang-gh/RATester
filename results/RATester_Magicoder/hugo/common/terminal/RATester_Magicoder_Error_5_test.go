package terminal

import (
	"fmt"
	"testing"
)

func TestError_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	expected := "\033[31mError: Test\033[0m"
	result := Error("Test")
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
