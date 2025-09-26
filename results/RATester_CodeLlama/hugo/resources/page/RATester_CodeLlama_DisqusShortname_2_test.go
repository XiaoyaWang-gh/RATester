package page

import (
	"fmt"
	"testing"
)

func TestDisqusShortname_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var testSite testSite
	var expected string

	// Act
	actual := testSite.DisqusShortname()

	// Assert
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
