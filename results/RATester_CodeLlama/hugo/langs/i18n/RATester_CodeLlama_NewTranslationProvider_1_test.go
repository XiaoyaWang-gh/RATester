package i18n

import (
	"fmt"
	"testing"
)

func TestNewTranslationProvider_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	translationProvider := NewTranslationProvider()
	// Assert
	if translationProvider == nil {
		t.Errorf("NewTranslationProvider() failed")
	}
}
