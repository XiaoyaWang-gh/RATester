package page

import (
	"fmt"
	"testing"
)

func TestTranslationBaseName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := nopPage(0)
	if p.TranslationBaseName() != "" {
		t.Errorf("Expected TranslationBaseName to be empty, but got %s", p.TranslationBaseName())
	}
}
