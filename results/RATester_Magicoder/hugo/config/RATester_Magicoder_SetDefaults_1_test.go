package config

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestSetDefaults_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	provider := &defaultConfigProvider{
		root: maps.Params{},
	}

	params := maps.Params{
		"key1": "value1",
		"key2": "value2",
	}

	provider.SetDefaults(params)

	if len(provider.root) != len(params) {
		t.Errorf("Expected provider root to have %d keys, but got %d", len(params), len(provider.root))
	}

	for k, v := range params {
		if provider.root[k] != v {
			t.Errorf("Expected provider root to have key %s with value %v, but got %v", k, v, provider.root[k])
		}
	}
}
