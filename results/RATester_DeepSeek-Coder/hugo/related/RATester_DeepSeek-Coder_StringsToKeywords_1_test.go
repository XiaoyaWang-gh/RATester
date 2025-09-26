package related

import (
	"fmt"
	"testing"
)

func TestStringsToKeywords_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cfg := IndexConfig{
		Name:                 "test",
		Type:                 "text",
		ApplyFilter:          true,
		Pattern:              "2006",
		Weight:               10,
		CardinalityThreshold: 50,
		ToLower:              true,
	}

	testStrings := []string{"test1", "test2", "test3"}
	keywords := cfg.StringsToKeywords(testStrings...)

	if len(keywords) != len(testStrings) {
		t.Errorf("Expected %d keywords, got %d", len(testStrings), len(keywords))
	}

	for i, kw := range keywords {
		if kw.String() != testStrings[i] {
			t.Errorf("Expected keyword %d to be %s, got %s", i, testStrings[i], kw.String())
		}
	}
}
