package assets

import (
	"fmt"
	"net/http"
	"testing"
)

func TestLoad_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	path := "./testdata"
	// Act
	Load(path)
	// Assert
	if FileSystem != http.Dir(path) {
		t.Errorf("FileSystem should be %v, but got %v", http.Dir(path), FileSystem)
	}
}
