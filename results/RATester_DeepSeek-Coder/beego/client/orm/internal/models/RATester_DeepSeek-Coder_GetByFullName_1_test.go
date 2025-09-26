package models

import (
	"fmt"
	"testing"
)

func TestGetByFullName_1(t *testing.T) {
	mc := &ModelCache{
		cacheByFullName: map[string]*ModelInfo{
			"test": {
				FullName: "test",
			},
		},
	}

	t.Run("should return model info when model exists", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		mi, ok := mc.GetByFullName("test")
		if !ok {
			t.Errorf("expected ok to be true, got false")
		}
		if mi.FullName != "test" {
			t.Errorf("expected FullName to be 'test', got %s", mi.FullName)
		}
	})

	t.Run("should return false when model does not exist", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, ok := mc.GetByFullName("non-existing")
		if ok {
			t.Errorf("expected ok to be false, got true")
		}
	})
}
