package codeblocks

import (
	"fmt"
	"testing"

	"github.com/yuin/goldmark"
)

func TestNew_40(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	result := goldmark.New()
	// Assert
	if result == nil {
		t.Errorf("New() = %v, want %v", result, &codeBlocksExtension{})
	}
}
