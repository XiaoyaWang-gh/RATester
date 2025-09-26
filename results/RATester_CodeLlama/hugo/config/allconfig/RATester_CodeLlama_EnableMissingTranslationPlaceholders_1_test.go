package allconfig

import (
	"fmt"
	"testing"
)

func TestEnableMissingTranslationPlaceholders_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := ConfigLanguage{}
	if c.EnableMissingTranslationPlaceholders() {
		t.Error("EnableMissingTranslationPlaceholders() should be false")
	}
}
