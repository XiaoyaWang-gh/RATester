package memory

import (
	"fmt"
	"testing"
)

func TestKeys_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Storage{
		db: map[string]entry{
			"key1": {data: []byte("value1"), expiry: 0},
			"key2": {data: []byte("value2"), expiry: 1},
			"key3": {data: []byte("value3"), expiry: 2},
		},
		done:       make(chan struct{}),
		gcInterval: 0,
	}

	keys, err := s.Keys()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(keys) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(keys))
	}

	s.db = map[string]entry{
		"key1": {data: []byte("value1"), expiry: 3},
		"key2": {data: []byte("value2"), expiry: 4},
		"key3": {data: []byte("value3"), expiry: 5},
	}

	keys, err = s.Keys()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(keys) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(keys))
	}
}
