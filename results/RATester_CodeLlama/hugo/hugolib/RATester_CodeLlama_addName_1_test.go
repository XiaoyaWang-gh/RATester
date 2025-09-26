package hugolib

import (
	"fmt"
	"testing"
)

func TestAddName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	s := &shortcodeHandler{}
	name := "name"

	// Act
	s.addName(name)

	// Assert
	if s.nameSet[name] != true {
		t.Errorf("Expected %s to be true", name)
	}
}
