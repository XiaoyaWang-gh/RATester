package filecache

import (
	"fmt"
	"testing"
)

func TestGetJSONCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var f Caches
	// Act
	result := f.GetJSONCache()
	// Assert
	if result == nil {
		t.Errorf("GetJSONCache() == nil, want not nil")
	}
}
