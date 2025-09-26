package filecache

import (
	"fmt"
	"testing"
)

func TestGetCSVCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var f Caches
	// Act
	result := f.GetCSVCache()
	// Assert
	if result == nil {
		t.Errorf("GetCSVCache() == nil, want not nil")
	}
}
