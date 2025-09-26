package web

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestGetMapData_1(t *testing.T) {
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
					MinTime:           1 * time.Second,
					MaxTime:           1 * time.Second,
					TotalTime:         1 * time.Second,
				},
			},
		},
	}

	result := urlMap.GetMapData()

	if len(result) != 1 {
		t.Errorf("Expected 1 result, got %d", len(result))
	}

	expected := map[string]interface{}{
		"request_url": "http://example.com",
		"method":      "GET",
		"times":       int64(1),
		"total_time":  "1s",
		"max_time":    "1s",
		"min_time":    "1s",
		"avg_time":    "1s",
	}

	if !reflect.DeepEqual(result[0], expected) {
		t.Errorf("Expected %v, got %v", expected, result[0])
	}
}
