package plugins

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPpSymbols_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	result := ppSymbols()

	// Check if the result is a map
	if reflect.TypeOf(result).Kind() != reflect.Map {
		t.Errorf("Expected result to be a map, but got %T", result)
	}

	// Check if the map has the expected keys
	expectedKeys := []string{"github.com/traefik/traefik/v3/pkg/plugins/plugins"}
	for _, key := range expectedKeys {
		if _, ok := result[key]; !ok {
			t.Errorf("Expected key %s not found in result", key)
		}
	}

	// Check if the map values are maps
	for _, innerMap := range result {
		if reflect.TypeOf(innerMap).Kind() != reflect.Map {
			t.Errorf("Expected inner map to be a map, but got %T", innerMap)
		}

		// Check if the inner map has the expected keys
		expectedInnerKeys := []string{"PP", "_PP"}
		for _, key := range expectedInnerKeys {
			if _, ok := innerMap[key]; !ok {
				t.Errorf("Expected inner key %s not found in inner map", key)
			}
		}
	}
}
