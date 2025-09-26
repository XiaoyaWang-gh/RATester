package allconfig

import (
	"fmt"
	"testing"
)

func TestFastRenderMode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	config := &Config{
		Internal: InternalConfig{
			FastRenderMode: true,
		},
	}
	configLang := ConfigLanguage{
		config: config,
	}

	result := configLang.FastRenderMode()

	if !result {
		t.Errorf("Expected FastRenderMode to be true, but got false")
	}
}
