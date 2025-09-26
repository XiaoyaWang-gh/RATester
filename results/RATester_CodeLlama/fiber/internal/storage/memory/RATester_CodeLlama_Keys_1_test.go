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
			"key1": {
				data:   []byte("value1"),
				expiry: 0,
			},
			"key2": {
				data:   []byte("value2"),
				expiry: 1000,
			},
			"key3": {
				data:   []byte("value3"),
				expiry: 1000,
			},
		},
	}

	keys, err := s.Keys()
	if err != nil {
		t.Fatal(err)
	}

	if len(keys) != 2 {
		t.Fatalf("expected 2 keys, got %d", len(keys))
	}

	if string(keys[0]) != "key1" {
		t.Fatalf("expected key1, got %s", string(keys[0]))
	}

	if string(keys[1]) != "key2" {
		t.Fatalf("expected key2, got %s", string(keys[1]))
	}
}
