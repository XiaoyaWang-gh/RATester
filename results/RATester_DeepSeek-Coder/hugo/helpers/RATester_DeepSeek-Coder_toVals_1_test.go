package helpers

import (
	"fmt"
	"testing"
)

func TestToVals_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	stats := &ProcessingStats{
		Name:            "Test",
		Pages:           10,
		PaginatorPages:  5,
		Static:          3,
		ProcessedImages: 2,
		Files:           1,
		Aliases:         1,
		Cleaned:         1,
	}

	vals := stats.toVals()

	if len(vals) != 8 {
		t.Errorf("Expected 8 values, got %d", len(vals))
	}

	expected := []string{"Pages", "Paginator pages", "Non-page files", "Static files", "Processed images", "Aliases", "Cleaned"}
	for i, v := range vals {
		if v.name != expected[i] {
			t.Errorf("Expected name %s, got %s", expected[i], v.name)
		}
	}
}
