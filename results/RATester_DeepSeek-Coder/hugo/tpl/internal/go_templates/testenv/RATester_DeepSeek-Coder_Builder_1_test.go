package testenv

import (
	"fmt"
	"os"
	"testing"
)

func TestBuilder_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	os.Setenv("GO_BUILDER_NAME", "Test Builder")
	result := Builder()
	if result != "Test Builder" {
		t.Errorf("Expected 'Test Builder', got '%s'", result)
	}
}
