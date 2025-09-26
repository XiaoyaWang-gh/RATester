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

	config := &Config{
		RootConfig: RootConfig{
			EnableMissingTranslationPlaceholders: true,
		},
	}

	configLanguage := ConfigLanguage{
		config: config,
	}

	result := configLanguage.EnableMissingTranslationPlaceholders()

	if !result {
		t.Error("Expected EnableMissingTranslationPlaceholders to return true, but got false")
	}
}
