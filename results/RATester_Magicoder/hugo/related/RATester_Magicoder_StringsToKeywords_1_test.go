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
		Type:                 "test",
		ApplyFilter:          true,
		Pattern:              "2006",
		Weight:               1,
		CardinalityThreshold: 50,
		ToLower:              true,
	}

	keywords := cfg.StringsToKeywords("test1", "test2", "test3")

	if len(keywords) != 3 {
		t.Errorf("Expected 3 keywords, got %d", len(keywords))
	}

	for _, kw := range keywords {
		if kw.String() != "test1" && kw.String() != "test2" && kw.String() != "test3" {
			t.Errorf("Expected keyword to be one of test1, test2, test3, got %s", kw.String())
		}
	}
}
