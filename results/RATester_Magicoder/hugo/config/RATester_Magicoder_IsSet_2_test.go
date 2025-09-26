package config

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestIsSet_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	provider := &defaultConfigProvider{
		root: maps.Params{
			"key1": "value1",
			"key2": "value2",
		},
	}

	// Test case 1: Key exists
	if !provider.IsSet("key1") {
		t.Error("Expected key1 to be set")
	}

	// Test case 2: Key does not exist
	if provider.IsSet("key3") {
		t.Error("Expected key3 to not be set")
	}

	// Test case 3: Key is case-insensitive
	if !provider.IsSet("KEY1") {
		t.Error("Expected KEY1 to be set")
	}
}
