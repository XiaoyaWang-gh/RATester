package allconfig

import (
	"fmt"
	"testing"
)

func TestNewContentEditor_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	config := &Config{
		RootConfig: RootConfig{
			NewContentEditor: "test_editor",
		},
	}

	configLanguage := ConfigLanguage{
		config: config,
	}

	expected := "test_editor"
	actual := configLanguage.NewContentEditor()

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
