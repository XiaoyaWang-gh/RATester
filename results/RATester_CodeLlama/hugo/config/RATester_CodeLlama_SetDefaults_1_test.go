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

	c := &defaultConfigProvider{
		root: make(maps.Params),
	}

	params := make(maps.Params)
	params["a"] = "1"
	params["b"] = "2"
	params["c"] = "3"

	c.SetDefaults(params)

	if len(c.root) != 3 {
		t.Errorf("Expected 3 params, got %d", len(c.root))
	}

	if c.root["a"] != "1" {
		t.Errorf("Expected a to be 1, got %s", c.root["a"])
	}

	if c.root["b"] != "2" {
		t.Errorf("Expected b to be 2, got %s", c.root["b"])
	}

	if c.root["c"] != "3" {
		t.Errorf("Expected c to be 3, got %s", c.root["c"])
	}
}
