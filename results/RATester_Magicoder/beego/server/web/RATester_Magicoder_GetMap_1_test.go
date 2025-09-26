package web

import (
	"fmt"
	"testing"
	"time"
)

func TestGetMap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	urlMap := &URLMap{
		urlmap: map[string]map[string]*Statistics{
			"http://example.com": {
				"GET": {
					RequestURL:        "http://example.com",
					RequestController: "example",
					RequestNum:        1,
					MinTime:           100 * time.Millisecond,
					MaxTime:           100 * time.Millisecond,
					TotalTime:         100 * time.Millisecond,
				},
			},
		},
	}

	result := urlMap.GetMap()

	if len(result) == 0 {
		t.Error("Expected non-empty result, but got empty")
	}

	if len(result["Fields"].([]string)) != 8 {
		t.Errorf("Expected 8 fields, but got %d", len(result["Fields"].([]string)))
	}

	if len(result["Data"].([][]string)) != 1 {
		t.Errorf("Expected 1 data, but got %d", len(result["Data"].([][]string)))
	}
}
