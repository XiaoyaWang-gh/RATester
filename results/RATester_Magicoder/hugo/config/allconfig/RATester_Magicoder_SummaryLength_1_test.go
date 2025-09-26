package allconfig

import (
	"fmt"
	"testing"
)

func TestSummaryLength_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	config := &Config{
		RootConfig: RootConfig{
			SummaryLength: 10,
		},
	}

	configLanguage := ConfigLanguage{
		config: config,
	}

	expected := 10
	actual := configLanguage.SummaryLength()

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
