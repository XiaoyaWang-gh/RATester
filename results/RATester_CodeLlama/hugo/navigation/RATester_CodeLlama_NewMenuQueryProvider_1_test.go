package navigation

import (
	"fmt"
	"testing"
)

func TestNewMenuQueryProvider_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var pagem PageMenusGetter
	var sitem MenusGetter
	var p Page

	// Act
	var result MenuQueryProvider = NewMenuQueryProvider(pagem, sitem, p)

	// Assert
	if result == nil {
		t.Errorf("NewMenuQueryProvider() = %v, want %v", result, nil)
	}
}
