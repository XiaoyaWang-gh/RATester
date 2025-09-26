package page

import (
	"fmt"
	"testing"
)

func TestTranslationKey_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := nopPage(0)
	if p.TranslationKey() != "" {
		t.Errorf("Expected TranslationKey to return an empty string, but got %s", p.TranslationKey())
	}
}
