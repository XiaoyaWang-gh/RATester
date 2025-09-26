package allconfig

import (
	"fmt"
	"testing"
)

func TestPrintUnusedTemplates_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	config := &Config{
		RootConfig: RootConfig{
			PrintUnusedTemplates: true,
		},
	}

	configLanguage := ConfigLanguage{
		config: config,
	}

	result := configLanguage.PrintUnusedTemplates()

	if result != true {
		t.Errorf("Expected true, but got %v", result)
	}
}
