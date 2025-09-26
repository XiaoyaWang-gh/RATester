package config

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestKeys_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	provider := &defaultConfigProvider{
		root: maps.Params{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		},
	}

	keys := provider.Keys()

	if len(keys) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(keys))
	}

	for _, key := range keys {
		if _, ok := provider.root[key]; !ok {
			t.Errorf("Key %s not found in root map", key)
		}
	}
}
