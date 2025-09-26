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

	tp := NewTranslationProvider()
	if tp == nil {
		t.Error("Expected a TranslationProvider, got nil")
	}
}
