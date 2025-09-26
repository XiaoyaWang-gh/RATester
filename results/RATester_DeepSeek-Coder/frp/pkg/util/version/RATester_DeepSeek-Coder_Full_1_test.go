package version

import (
	"fmt"
	"testing"
)

func TestFull_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	expected := "1.0.0"
	result := Full()
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
