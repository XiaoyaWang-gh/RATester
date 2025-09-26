package memory

import (
	"fmt"
	"testing"
)

func TestConn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	storage := &Storage{
		db: map[string]entry{
			"key1": {data: []byte("value1"), expiry: 0},
			"key2": {data: []byte("value2"), expiry: 0},
		},
	}

	result := storage.Conn()

	if len(result) != 2 {
		t.Errorf("Expected 2 entries, got %d", len(result))
	}

	if _, ok := result["key1"]; !ok {
		t.Error("Expected key1 to be in the result")
	}

	if _, ok := result["key2"]; !ok {
		t.Error("Expected key2 to be in the result")
	}
}
