package models

import (
	"fmt"
	"testing"
)

func TestBootstrap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mc := &ModelCache{
		cache:           make(map[string]*ModelInfo),
		cacheByFullName: make(map[string]*ModelInfo),
	}

	// Add some models to the cache
	mc.cache["model1"] = &ModelInfo{
		Name:     "Model1",
		FullName: "package.Model1",
		Fields: &Fields{
			Columns: map[string]*FieldInfo{
				"field1": {
					Rel: true,
				},
			},
		},
	}
	mc.cache["model2"] = &ModelInfo{
		Name:     "Model2",
		FullName: "package.Model2",
		Fields: &Fields{
			Columns: map[string]*FieldInfo{
				"field2": {
					Rel: true,
				},
			},
		},
	}

	// Call the method under test
	mc.Bootstrap()

	// Check if the method has been called correctly
	if !mc.done {
		t.Error("Bootstrap method not called")
	}

	// Add more tests to check the correctness of the method
}
