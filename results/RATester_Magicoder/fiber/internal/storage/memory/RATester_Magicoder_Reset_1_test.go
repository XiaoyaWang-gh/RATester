package memory

import (
	"fmt"
	"testing"
	"time"
)

func TestReset_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Storage{
		db: map[string]entry{
			"key1": {data: []byte("value1"), expiry: 123456789},
			"key2": {data: []byte("value2"), expiry: 987654321},
		},
		done:       make(chan struct{}),
		gcInterval: 1 * time.Second,
	}

	err := s.Reset()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(s.db) != 0 {
		t.Errorf("Expected db to be empty, got %v", s.db)
	}
}
