package config

import (
	"fmt"
	"testing"
)

func TestDetectLegacyINIFormatFromFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	path := "./testdata/legacy.ini"
	expected := true

	// Act
	actual := DetectLegacyINIFormatFromFile(path)

	// Assert
	if actual != expected {
		t.Errorf("DetectLegacyINIFormatFromFile() = %v, want %v", actual, expected)
	}
}
