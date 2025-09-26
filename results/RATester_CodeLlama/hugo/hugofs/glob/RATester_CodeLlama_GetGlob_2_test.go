package glob

import (
	"fmt"
	"testing"
)

func TestGetGlob_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	pattern := "*.go"

	// Act
	glob, err := GetGlob(pattern)

	// Assert
	if err != nil {
		t.Errorf("GetGlob() error = %v, want nil", err)
		return
	}

	if !glob.Match("main.go") {
		t.Errorf("GetGlob() = %v, want %v", glob.Match("main.go"), true)
	}
}
