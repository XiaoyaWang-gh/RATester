package hugo

import (
	"fmt"
	"testing"
)

func TestGetDependencyListNonGo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	// Assert
}
