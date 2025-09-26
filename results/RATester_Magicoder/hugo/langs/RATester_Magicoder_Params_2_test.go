package langs

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestParams_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &Language{
		params: maps.Params{
			"key1": "value1",
			"key2": "value2",
		},
	}

	result := l.Params()

	if len(result) != 2 {
		t.Errorf("Expected 2 items in the map, got %d", len(result))
	}

	if result["key1"] != "value1" {
		t.Errorf("Expected 'value1', got '%s'", result["key1"])
	}

	if result["key2"] != "value2" {
		t.Errorf("Expected 'value2', got '%s'", result["key2"])
	}
}
