package navigation

import (
	"fmt"
	"testing"
)

func TestIsMenuCurrent_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var m nopPageMenus
	var menuID string
	var inme *MenuEntry

	// Act
	var actual bool
	actual = m.IsMenuCurrent(menuID, inme)

	// Assert
	var expected bool
	expected = false
	if actual != expected {
		t.Errorf("actual %v != expected %v", actual, expected)
	}
}
