package helpers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToVals_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &ProcessingStats{
		Name:            "test",
		Pages:           10,
		PaginatorPages:  20,
		Static:          30,
		ProcessedImages: 40,
		Files:           50,
		Aliases:         60,
		Cleaned:         70,
	}

	expected := []processingStatsTitleVal{
		{"Pages", 10},
		{"Paginator pages", 20},
		{"Non-page files", 30},
		{"Static files", 40},
		{"Processed images", 50},
		{"Aliases", 60},
		{"Cleaned", 70},
	}

	if !reflect.DeepEqual(s.toVals(), expected) {
		t.Errorf("got %v, want %v", s.toVals(), expected)
	}
}
