package maps

import (
	"errors"
	"fmt"
	"testing"
)

func TestGetOrCreate_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cache := &Cache[string, int]{m: make(map[string]int)}

	key := "test"
	value := 123

	// Test creating a new value
	val, err := cache.GetOrCreate(key, func() (int, error) {
		return value, nil
	})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != value {
		t.Errorf("Expected %d, got %d", value, val)
	}

	// Test getting an existing value
	val, err = cache.GetOrCreate(key, func() (int, error) {
		t.Errorf("Unexpected call to create function")
		return 0, nil
	})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != value {
		t.Errorf("Expected %d, got %d", value, val)
	}

	// Test error from create function
	_, err = cache.GetOrCreate(key, func() (int, error) {
		return 0, errors.New("test error")
	})
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
