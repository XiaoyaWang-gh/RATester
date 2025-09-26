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

	expected := "\033[31mError: test\033[0m"
	result := Error("test")
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
