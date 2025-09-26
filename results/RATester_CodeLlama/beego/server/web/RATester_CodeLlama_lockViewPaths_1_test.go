package web

import (
	"fmt"
	"testing"
)

func TestLockViewPaths_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	lockViewPaths()
	// Assert
	if beeViewPathTemplateLocked != true {
		t.Errorf("beeViewPathTemplateLocked should be true")
	}
}
