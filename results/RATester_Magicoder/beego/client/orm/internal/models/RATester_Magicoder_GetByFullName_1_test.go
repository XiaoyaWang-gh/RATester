package models

import (
	"fmt"
	"testing"
)

func TestGetByFullName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mc := &ModelCache{
		cacheByFullName: map[string]*ModelInfo{
			"test": {
				FullName: "test",
			},
		},
	}

	mi, ok := mc.GetByFullName("test")
	if !ok {
		t.Fatal("Expected to find model info")
	}

	if mi.FullName != "test" {
		t.Fatalf("Expected model info to have full name 'test', got '%s'", mi.FullName)
	}
}
