package helpers

import (
	"fmt"
	"testing"
)

func TesttoVals_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	stats := ProcessingStats{
		Name:            "TestStats",
		Pages:           10,
		PaginatorPages:  5,
		Static:          2,
		ProcessedImages: 1,
		Files:           15,
		Aliases:         3,
		Cleaned:         2,
	}

	vals := stats.toVals()

	if len(vals) != 7 {
		t.Errorf("Expected 7 values, got %d", len(vals))
	}

	expected := []processingStatsTitleVal{
		{"Pages", 10},
		{"Paginator pages", 5},
		{"Non-page files", 15},
		{"Static files", 2},
		{"Processed images", 1},
		{"Aliases", 3},
		{"Cleaned", 2},
	}

	for i, val := range vals {
		if val.name != expected[i].name || val.val != expected[i].val {
			t.Errorf("Expected %v, got %v", expected[i], val)
		}
	}
}
